package user_dal

import (
	"github.com/alf-grindel/dawn/pkg/errno"
	"log"
)

func (s *UserDal) GetUserByAccount(userAccount string) (*User, error) {
	user := &User{}

	if err := s.db.Where("user_account = ?", userAccount).First(&user).Error; err != nil {
		log.Printf("dal.query: can not find the user by acount, %s\n", err)
		return nil, errno.SystemErr
	}
	return user, nil
}
