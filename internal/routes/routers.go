package routes

import (
	"github.com/alf-grindel/dawn/internal/app"
	"github.com/alf-grindel/dawn/internal/middleware"
	"github.com/alf-grindel/dawn/pkg/constants"
	"github.com/gorilla/mux"
)

func SetUpRouters(app *app.Application) *mux.Router {
	r := mux.NewRouter()

	publicR := r.PathPrefix("/user").Subrouter()
	{
		publicR.HandleFunc("/register", app.UserHandler.Register).Methods("POST")
		publicR.HandleFunc("/logout", app.UserHandler.Logout).Methods("POST")
		publicR.HandleFunc("/login", app.UserHandler.Login).Methods("POST")
		publicR.HandleFunc("/delete", app.UserHandler.Delete).Methods("POST")

		publicR.HandleFunc("/update", app.UserHandler.Update).Methods("PUT")

		publicR.HandleFunc("/get/login", app.UserHandler.GetLoginUser).Methods("GET")
	}

	// admin
	authR := r.PathPrefix("/admin").Subrouter()
	{
		authR.Use(middleware.AuthMiddleware(constants.Store))

		authR.HandleFunc("/add", app.AdminHandler.Add).Methods("POST")
		authR.HandleFunc("/delete", app.AdminHandler.Delete).Methods("POST")

		authR.HandleFunc("/update", app.AdminHandler.Update).Methods("PUT")

		authR.HandleFunc("/get/user", app.AdminHandler.Get).Methods("GET")
		authR.HandleFunc("/get/list/user", app.AdminHandler.GetList).Methods("GET")
	}

	return r
}
