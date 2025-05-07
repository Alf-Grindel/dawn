package user_dal

import (
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/pkg/errno"
	"log"
)

func (s *UserDal) CreateUser(userAccount, hashPassword string) (int64, error) {
	user := &data.User{
		Id:           s.Snowflake.GenerateID(),
		Account:      userAccount,
		UserPassword: hashPassword,
	}

	if err := s.db.Select("id", "account", "user_password").Create(&user).Error; err != nil {
		log.Printf("dal.create: create user failed, %s", err)
		return -1, errno.SystemErr
	}
	return user.Id, nil
}

func (s *UserDal) Add(user *data.User) (int64, error) {
	if err := s.db.
		Select("id", "account", "user_password", "user_name", "user_avatar", "user_profile", "user_role").
		Create(&user).Error; err != nil {
		log.Printf("dal.create: add user failed, %s", err)
		return -1, errno.SystemErr
	}
	return user.Id, nil
}
