package di

import (
	"context"
	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"github.com/haseebh/weatherapp_auth/internal/infrastrcuture/datastore"
	"github.com/haseebh/weatherapp_auth/pkg/config"
)

const UserCollection = "users"

func GetUserRepository() repository.UserRepository {
	return datastore.NewUserDB(GetBaseDatabase(), UserCollection)
}

func GetBaseDatabase() *datastore.Database {
	return datastore.NewDatabase(context.Background(), config.LoadConfig())
}
