package main

import (
	"github.com/haseebh/weatherapp_temperature/internal/di"
	"github.com/haseebh/weatherapp_temperature/internal/handlers"
	"github.com/haseebh/weatherapp_temperature/internal/utils"
	"github.com/haseebh/weatherapp_temperature/pkg/server"
	"log"
)

func init() {

}
func main() {
	srv := server.GetHTTPServer()
	srv.TemperatureHandlers(
		handlers.NewTemperatureHandler(
			di.GetTemperatureUseCase(),
		))

	go func() {
		err := di.GetMessageQueueRepository().Consume()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go utils.FetchWeatherData()
	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
