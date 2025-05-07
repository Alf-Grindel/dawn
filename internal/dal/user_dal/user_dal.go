package user_dal

import (
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/internal/model/request"
	"github.com/alf-grindel/dawn/pkg/utils"
	"gorm.io/gorm"
)

type UserDal struct {
	db        *gorm.DB
	Snowflake *utils.Snowflake
}

func NewUserDal(db *gorm.DB, snowflake *utils.Snowflake) *UserDal {
	return &UserDal{
		db:        db,
		Snowflake: snowflake,
	}
}

type UserStore interface {
	CreateUser(userAccount, hashPassword string) (int64, error)
	GetUserByAccount(userAccount string) (*data.User, error)
	UpdateUser(user *data.User) (*data.User, error)
	GetUserById(id int64) (*data.User, error)
	GetUserByRole(role string) (*data.User, error)
	Add(user *data.User) (int64, error)
	GetUserSingle(req *request.UserQueryRequest) (*data.User, error)
	GetUserList(req *request.UserQueryListRequest) ([]*data.User, int64, error)
	DeleteUser(id int64) (bool, error)
}
