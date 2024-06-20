package di

import (
	"context"
	"github.com/haseebh/weatherapp_temperature/internal/entities/repository"
	"github.com/haseebh/weatherapp_temperature/internal/infrastrcuture/datastore"
	"github.com/haseebh/weatherapp_temperature/pkg/config"
)

const TemperatureCollection = "temperature"

func GetTemperatureRepository() repository.TemperatureRepository {
	return datastore.NewTemperatureDB(GetBaseDatabase(), TemperatureCollection)
}

func GetBaseDatabase() *datastore.Database {
	return datastore.NewDatabase(context.Background(), config.LoadConfig())
}
