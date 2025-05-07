package user_dal

import (
	"github.com/alf-grindel/dawn/pkg/errno"
	"log"
)

func (s *UserDal) CreateUser(userAccount, hashPassword string) (int64, error) {
	user := &User{
		Id:           s.snowflake.GenerateID(),
		Account:      userAccount,
		UserPassword: hashPassword,
	}

	if err := s.db.Select("id", "account", "user_password").Create(&user).Error; err != nil {
		log.Printf("dal.create: create user failed, %s", err)
		return -1, errno.SystemErr
	}
	return user.Id, nil
}
