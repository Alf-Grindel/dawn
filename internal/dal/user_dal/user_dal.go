package user_dal

import (
	"github.com/alf-grindel/dawn/pkg/utils"
	"gorm.io/gorm"
)

type UserDal struct {
	db        *gorm.DB
	snowflake *utils.Snowflake
}

func NewUserDal(db *gorm.DB, snowflake *utils.Snowflake) *UserDal {
	return &UserDal{
		db:        db,
		snowflake: snowflake,
	}
}

type UserStore interface {
	CreateUser(userAccount, hashPassword string) (int64, error)
	GetUserByAccount(userAccount string) (*User, error)
	UpdateUser(*User) (*User, error)
}
