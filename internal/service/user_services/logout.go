package user_services

import (
	"github.com/alf-grindel/dawn/pkg/constants"
	"net/http"
)

func (s *UserService) Logout(rw http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, constants.UserLoginState)
	if session.Values["login"] != true {
		s.logger.Println("service.logout: no login user")
		return false
	}
	session.Values["user"] = nil
	session.Values["login"] = false
	err := session.Save(r, rw)
	if err != nil {
		s.logger.Printf("service.logout: session保存失败, %s", err)
		return false
	}
	return true
}
