package user_dal

import (
	"github.com/alf-grindel/dawn/pkg/errno"
	"log"
)

func (s *UserDal) UpdateUser(user *User) (*User, error) {
	if err := s.db.Model(&User{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		log.Printf("dal.update: updates user failed, %s\n", err)
		return nil, errno.SystemErr
	}
	return user, nil
}
