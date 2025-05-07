package user_services

import (
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/internal/model/request"
	"github.com/alf-grindel/dawn/internal/model/safe"
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/utils"
)

func (s *UserService) Update(user *request.UpdateRequest) (*safe.User, error) {
	if user == nil {
		return nil, errno.ParamsErr
	}
	if user.UserPassword != "" {
		if len(user.UserPassword) < 8 {
			return nil, errno.ParamsErr.WithMessage("用户密码过短")
		}
		user.UserPassword = utils.HashPassword(user.UserPassword)
	}

	u := &data.User{
		Id:           user.Id,
		UserPassword: user.UserPassword,
		UserName:     user.UserName,
		UserAvatar:   user.UserAvatar,
		UserProfile:  user.UserProfile,
	}

	currentUser, err := s.userDal.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	return GetUserSafe(currentUser), nil
}
