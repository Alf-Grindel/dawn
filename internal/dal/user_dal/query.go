package user_dal

import (
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/internal/model/request"
	"github.com/alf-grindel/dawn/pkg/errno"
	"log"
)

func (s *UserDal) GetUserByAccount(userAccount string) (*data.User, error) {
	user := &data.User{}

	if err := s.db.Where("user_account = ?", userAccount).First(&user).Error; err != nil {
		log.Printf("dal.query: can not find the user by acount, %s\n", err)
		return nil, errno.SystemErr
	}
	return user, nil
}

func (s *UserDal) GetUserById(id int64) (*data.User, error) {
	user := &data.User{}

	if err := s.db.Where("id = ? and is_delete = 0", id).First(&user).Error; err != nil {
		log.Printf("dal.query: can not find the user by id, %s\n", err)
		return nil, errno.SystemErr
	}
	return user, nil
}
func (s *UserDal) GetUserByRole(role string) (*data.User, error) {
	user := &data.User{}

	if err := s.db.Where("user_role = ? and is_delete = 0", role).First(&user).Error; err != nil {
		log.Printf("dal.query: can not find the user by userRole, %s\n", err)
		return nil, errno.SystemErr
	}
	return user, nil
}

func (s *UserDal) GetUserSingle(req *request.UserQueryRequest) (*data.User, error) {
	user := &data.User{}
	query := s.db.Model(&data.User{}).Where("is_delete = 0")
	if req.Id != -1 {
		query = query.Where("id = ?", req.Id)
	}
	if req.UserAccount != "" {
		query = query.Where("user_account = ?", req.UserAccount)
	}
	if req.UserRole != "" {
		query = query.Where("user_role = ?", req.UserRole)
	}
	if err := query.First(&user).Error; err != nil {
		log.Printf("dal.query: can not find the user, %s\n", err)
		return nil, errno.SystemErr
	}
	return user, nil
}

func (s *UserDal) GetUserList(req *request.UserQueryListRequest) ([]*data.User, int64, error) {
	users := []*data.User{}
	total := int64(0)

	query := s.db.Model(&data.User{}).Where("is_delete = 0")
	if req.Id != -1 {
		query = query.Where("id = ?", req.Id)
	}
	if req.UserAccount != "" {
		query = query.Where("user_account = ?", req.UserAccount)
	}
	if req.UserRole != "" {
		query = query.Where("user_role = ?", req.UserRole)
	}

	if err := query.Count(&total).Error; err != nil {
		log.Printf("dal.query: can not count the user, %s\n", err)
		return nil, -1, errno.SystemErr
	}

	offset := (req.Current - 1) * req.PageSize
	if err := query.Limit(int(req.PageSize)).Offset(int(offset)).Find(&users).Error; err != nil {
		log.Printf("dal.query: can not find the user, %s\n", err)
		return nil, -1, errno.SystemErr
	}
	return users, total, nil
}
