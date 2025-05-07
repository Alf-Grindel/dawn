package api

import (
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/internal/model/request"
	"github.com/alf-grindel/dawn/internal/service/user_services"
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/resp"
	"github.com/alf-grindel/dawn/pkg/utils"
	"log"
	"net/http"
)

type UserHandler struct {
	userDal *user_dal.UserDal
	logger  *log.Logger
}

func NewUserHandler(userDal *user_dal.UserDal, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userDal: userDal,
		logger:  logger,
	}
}
func (h *UserHandler) Register(rw http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.user_handler: decoding register request failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	id, err := user_services.NewUserService(r.Context(), h.logger, h.userDal).Register(req.UserAccount, req.UserPassword, req.CheckPassword)
	if err != nil {
		h.logger.Printf("[ERROR]api.user_handler: register user failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"id": id})
}

func (h *UserHandler) Login(rw http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.user_handler: decoding login request failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	user, err := user_services.NewUserService(r.Context(), h.logger, h.userDal).Login(req.UserAccount, req.UserPassword, rw, r)
	if err != nil {
		h.logger.Printf("[ERROR]api.user_handler: login user failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"user": user})
}

func (h *UserHandler) GetLoginUser(rw http.ResponseWriter, r *http.Request) {
	currentUser := user_services.NewUserService(r.Context(), h.logger, h.userDal).GetLoginUser(r)
	user := user_services.GetUserSafe(currentUser)
	if user == nil {
		h.logger.Println("[ERROR]api.user_handler: failed to get login user")
		resp.WriteJson(rw, errno.NotLoginErr, nil)
		return
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"user": user})
}

func (h *UserHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	isLogout := user_services.NewUserService(r.Context(), h.logger, h.userDal).Logout(rw, r)
	if !isLogout {
		h.logger.Println("[ERROR]api.user_handler: failed logout user")
		resp.WriteJson(rw, errno.ParamsErr, nil)
		return
	}
	resp.WriteJson(rw, errno.Success, nil)
}

func (h *UserHandler) Update(rw http.ResponseWriter, r *http.Request) {
	var req request.UpdateRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.user_handler: decoding update request failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	currentUser := user_services.NewUserService(r.Context(), h.logger, h.userDal).GetLoginUser(r)
	if currentUser == nil {
		resp.WriteJson(rw, errno.NotLoginErr, nil)
		return
	}
	if currentUser.Id != req.Id {
		resp.WriteJson(rw, errno.NoAuthErr, nil)
		return
	}
	user, err := user_services.NewUserService(r.Context(), h.logger, h.userDal).Update(&req)
	if err != nil {
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"user": user})
}

func (h *UserHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	currentUser := user_services.NewUserService(r.Context(), h.logger, h.userDal).GetLoginUser(r)
	if currentUser == nil {
		resp.WriteJson(rw, errno.NotLoginErr, nil)
		return
	}

	if _, err := h.userDal.DeleteUser(currentUser.Id); err != nil {
		h.logger.Printf("[ERROR]api.user_handler: delete user failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	resp.WriteJson(rw, errno.Success, nil)
}
