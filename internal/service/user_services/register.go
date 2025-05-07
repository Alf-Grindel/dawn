package user_services

import (
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/utils"
	"regexp"
)

func (s *UserService) Register(userAccount, userPassword, checkPassword string) (int64, error) {
	if len(userAccount) == 0 || len(userPassword) == 0 || len(checkPassword) == 0 {
		return -1, errno.ParamsErr.WithMessage("参数为空")
	}
	if len(userAccount) < 4 {
		return -1, errno.ParamsErr.WithMessage("用户账户过短")
	}
	if len(userPassword) < 8 {
		return -1, errno.ParamsErr.WithMessage("用户密码过短")
	}
	if userPassword != checkPassword {
		return -1, errno.ParamsErr.WithMessage("两次输入不一致")
	}
	reg := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !reg.MatchString(userAccount) {
		return -1, errno.ParamsErr.WithMessage("账号存在未知字符")
	}

	if u, _ := s.userDal.GetUserByAccount(userAccount); u != nil {
		return -1, errno.ParamsErr.WithMessage("账号重复")
	}
	hashPassword := utils.HashPassword(userPassword)
	id, err := s.userDal.CreateUser(userAccount, hashPassword)
	if err != nil {
		return -1, err
	}
	return id, nil
}
