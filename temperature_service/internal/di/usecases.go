package di

import "github.com/haseebh/weatherapp_temperature/internal/usecases"

func GetTemperatureUseCase() usecases.TemperatureUseCase {
	return usecases.NewTemperatureUseCase(GetTemperatureRepository())
}
