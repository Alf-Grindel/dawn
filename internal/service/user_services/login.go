package user_services

import (
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/internal/model/safe"
	"github.com/alf-grindel/dawn/pkg/constants"
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/utils"
	"net/http"
	"regexp"
)

func (s *UserService) Login(userAccount, userPassword string, rw http.ResponseWriter, r *http.Request) (*safe.User, error) {
	if len(userAccount) == 0 || len(userPassword) == 0 {
		return nil, errno.ParamsErr.WithMessage("参数为空")
	}
	if len(userAccount) < 4 {
		return nil, errno.ParamsErr.WithMessage("用户账户过短")
	}
	if len(userPassword) < 8 {
		return nil, errno.ParamsErr.WithMessage("用户密码过短")
	}
	reg := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !reg.MatchString(userAccount) {
		return nil, errno.ParamsErr.WithMessage("账号存在未知字符")
	}
	u, err := s.userDal.GetUserByAccount(userAccount)
	if err != nil {
		return nil, errno.SystemErr.WithMessage("用户不存在或密码错误")
	}
	if !utils.ComparePassword(userPassword, u.UserPassword) {
		return nil, errno.SystemErr.WithMessage("用户不存在或密码错误")
	}
	session, _ := constants.Store.Get(r, constants.UserLoginState)
	session.Values["user"] = u
	session.Values["login"] = true
	err = session.Save(r, rw)
	if err != nil {
		s.logger.Println("session保存失败: %s", err)
		return nil, errno.SystemErr
	}
	return GetUserSafe(u), nil
}

func (s *UserService) GetLoginUser(r *http.Request) *data.User {
	session, _ := constants.Store.Get(r, constants.UserLoginState)
	if session.Values["login"] != true {
		s.logger.Println("service.login: no login user")
		return nil
	}
	u, ok := session.Values["user"].(data.User)
	if !ok {
		s.logger.Println("service.login: session user type mismatch or missing")
		return nil
	}
	return &u
}
