package app

import (
	"github.com/alf-grindel/dawn/internal/api"
	"github.com/alf-grindel/dawn/internal/dal"
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/pkg/utils"
	"log"
	"os"
)

type Application struct {
	Logger       *log.Logger
	UserHandler  *api.UserHandler
	AdminHandler *api.AdminHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "dawn", log.LstdFlags)
	db := dal.Open()
	snowflake := utils.NewSnowflake(0)
	userDal := user_dal.NewUserDal(db, snowflake)

	userHandler := api.NewUserHandler(userDal, logger)
	adminHandler := api.NewAdminHandler(userDal, logger)

	app := &Application{
		Logger:       logger,
		UserHandler:  userHandler,
		AdminHandler: adminHandler,
	}

	return app, nil
}
