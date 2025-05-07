package main

import (
	"context"
	"encoding/gob"
	"github.com/alf-grindel/dawn/conf"
	. "github.com/alf-grindel/dawn/internal/app"
	"github.com/alf-grindel/dawn/internal/dal"
	"github.com/alf-grindel/dawn/internal/dal/user_dal"
	"github.com/alf-grindel/dawn/internal/routes"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	conf.Init()
	dal.Init()
	gob.Register(user_dal.User{})
}

func main() {
	app, err := NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetUpRouters(app)

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(), // allow cookie
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
		handlers.ExposedHeaders([]string{"*"}),
	)

	server := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      corsOptions(r),    // set the default handler
		ErrorLog:     app.Logger,        // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connecting using TCP Keep-Active
	}

	go func() {
		app.Logger.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			app.Logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	app.Logger.Println("Got signal: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(ctx)
}
