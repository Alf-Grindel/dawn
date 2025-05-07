package user_services

import (
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/internal/service"
	"github.com/alf-grindel/dawn/pkg/constants"
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/utils"
	"github.com/gorilla/sessions"
	"net/http"
	"regexp"
)

var store = sessions.NewCookieStore([]byte("dawn0814."))

func (s *UserService) Login(userAccount, userPassword string, rw http.ResponseWriter, r *http.Request) (*service.User, error) {
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
	session, _ := store.Get(r, constants.UserLoginState)
	session.Values["user"] = u
	session.Values["login"] = true
	err = session.Save(r, rw)
	if err != nil {
		s.logger.Println("session保存失败: %s", err)
		return nil, errno.SystemErr
	}
	return s.GetLoginUserSafe(u), nil
}

func (s *UserService) GetLoginUser(r *http.Request) *user_dal.User {
	session, _ := store.Get(r, constants.UserLoginState)
	if session.Values["login"] != true {
		s.logger.Println("service.login: no login user")
		return nil
	}
	u, ok := session.Values["user"].(user_dal.User)
	if !ok {
		s.logger.Println("service.login: session user type mismatch or missing")
		return nil
	}
	return &u
}

func (s *UserService) GetLoginUserSafe(user *user_dal.User) *service.User {
	if user == nil {
		return nil
	}
	u := &service.User{
		UserAccount: user.Account,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		UserProfile: user.UserProfile,
		UserRole:    user.UserRole,
		CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  user.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return u
}
