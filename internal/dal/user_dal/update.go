package user_dal

import (
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/pkg/errno"
	"log"
)

func (s *UserDal) UpdateUser(user *data.User) (*data.User, error) {
	if err := s.db.Model(&data.User{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		log.Printf("dal.update: updates user failed, %s\n", err)
		return nil, errno.SystemErr
	}
	u, err := s.GetUserById(user.Id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserDal) DeleteUser(id int64) (bool, error) {
	if err := s.db.Model(&data.User{}).Where("id = ?", id).Update("is_delete", 1).Error; err != nil {
		log.Printf("dal.update: delete user failed, %s\n", err)
		return false, errno.SystemErr
	}
	return true, nil
}
