package main

import (
	"github.com/haseebh/weatherapp_auth/internal/di"
	"github.com/haseebh/weatherapp_auth/internal/handlers"
	"github.com/haseebh/weatherapp_auth/pkg/server"
	"log"
)

func main() {
	srv := server.GetHTTPServer()

	srv.UserHandlers(
		handlers.NewUserHandler(
			di.GetUserUseCase()))
	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
