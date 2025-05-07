package api

import (
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/internal/model/request"
	"github.com/alf-grindel/dawn/internal/model/safe"
	"github.com/alf-grindel/dawn/internal/service/user_services"
	"github.com/alf-grindel/dawn/pkg/constants"
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/resp"
	"github.com/alf-grindel/dawn/pkg/utils"
	"log"
	"net/http"
	"strings"
)

type AdminHandler struct {
	adminDal *user_dal.UserDal
	logger   *log.Logger
}

func NewAdminHandler(adminDal *user_dal.UserDal, logger *log.Logger) *AdminHandler {
	return &AdminHandler{
		adminDal: adminDal,
		logger:   logger,
	}
}

var DefaultPassword = "12345678"

func (h *AdminHandler) Add(rw http.ResponseWriter, r *http.Request) {
	var req request.UserAddRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: decoding add request failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	password := utils.HashPassword(DefaultPassword)
	user := data.User{
		Account:      req.UserAccount,
		UserPassword: password,
		UserName:     strings.TrimSpace(req.UserName),
		UserAvatar:   strings.TrimSpace(req.UserAvatar),
		UserProfile:  strings.TrimSpace(req.UserProfile),
		UserRole:     req.UserRole,
	}
	if user.UserRole == "" {
		user.UserRole = "user"
	}
	id, err := h.adminDal.Add(&user)
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: add user failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"id": id})
}

func (h *AdminHandler) Get(rw http.ResponseWriter, r *http.Request) {
	var req request.UserQueryRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: decoding get user failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	user, err := h.adminDal.GetUserSingle(&req)
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: get user failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	u := user_services.GetUserSafe(user)
	if u == nil {
		h.logger.Printf("[ERROR]api.admin_handler: get user failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"user": u})
}

func (h *AdminHandler) GetList(rw http.ResponseWriter, r *http.Request) {
	var req request.UserQueryListRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: decoding get users request failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	if req.Current <= 0 {
		req.Current = constants.Current
	}
	if req.PageSize <= 0 {
		req.PageSize = constants.PageSize
	}
	u, total, err := h.adminDal.GetUserList(&req)
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: get users failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	us := user_services.GetUserSafeList(u)
	if us == nil {
		h.logger.Printf("[ERROR]api.admin_handler: get users failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	user := safe.Page[*safe.User]{
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    total,
		Data:     us,
	}
	resp.WriteJson(rw, errno.Success, resp.Data{"data": user})
}

func (h *AdminHandler) Update(rw http.ResponseWriter, r *http.Request) {
	var req request.UserUpdateRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: decoding update failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	user := data.User{
		Id:          req.Id,
		UserName:    req.UserName,
		UserAvatar:  req.UserAvatar,
		UserProfile: req.UserProfile,
		UserRole:    req.UserRole,
	}
	u, err := h.adminDal.UpdateUser(&user)
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: update user failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}
	us := user_services.GetUserSafe(u)
	resp.WriteJson(rw, errno.Success, resp.Data{"user": us})
}

func (h *AdminHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	var req request.UserDeleteRequest
	err := utils.FromJSON(&req, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: decoding delete failed, %s\n", err)
		resp.WriteJson(rw, errno.SystemErr, nil)
		return
	}

	if _, err := h.adminDal.DeleteUser(req.Id); err != nil {
		h.logger.Printf("[ERROR]api.admin_handler: delete user failed, %s\n", err)
		resp.WriteJson(rw, errno.ConvertErr(err), nil)
		return
	}
	resp.WriteJson(rw, errno.Success, nil)
}
