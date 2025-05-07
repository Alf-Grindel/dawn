package routes

import (
	"github.com/alf-grindel/dawn/internal/app"
	"github.com/gorilla/mux"
)

func SetUpRouters(app *app.Application) *mux.Router {
	r := mux.NewRouter()

	postR := r.Methods("POST").Subrouter()
	postR.HandleFunc("/user/register", app.UserHandler.Register)
	postR.HandleFunc("/user/logout", app.UserHandler.Logout)

	getR := r.Methods("GET").Subrouter()
	getR.HandleFunc("/user/login", app.UserHandler.Login)
	getR.HandleFunc("/user/login/get", app.UserHandler.GetLoginUser)

	return r
}
