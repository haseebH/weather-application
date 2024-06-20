package utils

import (
	"context"
	"fmt"
	"github.com/haseebh/weatherapp_temperature/internal/di"
	"time"
)

var (
	locations = []string{"London"}
)

func FetchWeatherData() {
	endTime := time.Now().UTC()

	sTime := time.Date(endTime.Year(), endTime.Month()-1, endTime.Day(), 0, 0, 0, 0, time.UTC)
	//todo: uncomment this one for 3 years of data
	//sTime := time.Date(endTime.Year()-3, endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.UTC)

	for i := range locations {
		err := di.GetTemperatureUseCase().FetchAndStoreTemperature(context.Background(), locations[i], sTime, endTime)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("data fetched successfully")
}
