package constants

import "github.com/gorilla/sessions"

const (
	UserLoginState = "user_login"
	DsnTemplate    = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	SALT           = "Dawn0814"

	// page list
	Current  = 1
	PageSize = 10
)

var Store = sessions.NewCookieStore([]byte("dawn0814."))
