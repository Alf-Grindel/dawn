package user_services

import (
	"context"
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/internal/service"
	"log"
	"net/http"
)

type UserService struct {
	ctx     context.Context
	logger  *log.Logger
	userDal *user_dal.UserDal
}

func NewUserService(ctx context.Context, logger *log.Logger, userDal *user_dal.UserDal) *UserService {
	return &UserService{
		ctx:     ctx,
		logger:  logger,
		userDal: userDal,
	}
}

type UserServices interface {
	Register(userAccount, userPassword, checkPassword string) (int64, error)
	Login(userAccount, userPassword string, rw http.ResponseWriter, r *http.Request) (*service.User, error)
	GetLoginUser(r *http.Request) *user_dal.User
	GetLoginUserSafe(user *user_dal.User) *service.User
	Logout(rw http.ResponseWriter, r *http.Request) bool
}
