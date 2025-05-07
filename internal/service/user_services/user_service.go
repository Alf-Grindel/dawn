package user_services

import (
	"context"
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/internal/model/request"
	"github.com/alf-grindel/dawn/internal/model/safe"
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
	Login(userAccount, userPassword string, rw http.ResponseWriter, r *http.Request) (*safe.User, error)
	GetLoginUser(r *http.Request) *data.User
	Logout(rw http.ResponseWriter, r *http.Request) bool
	Update(user *request.UpdateRequest) (*safe.User, error)
}

func GetUserSafe(user *data.User) *safe.User {
	if user == nil {
		return nil
	}
	u := &safe.User{
		Id:          user.Id,
		UserAccount: user.Account,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		UserProfile: user.UserProfile,
		UserRole:    user.UserRole,
		CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  user.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return u
}

func GetUserSafeList(users []*data.User) []*safe.User {
	u := []*safe.User{}
	if users == nil {
		return nil
	}
	for _, user := range users {
		u = append(u, GetUserSafe(user))
	}
	return u
}
