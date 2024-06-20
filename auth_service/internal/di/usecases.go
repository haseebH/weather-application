package di

import "github.com/haseebh/weatherapp_auth/internal/usecases"

func GetUserUseCase() usecases.UserUseCase {
	return usecases.NewUserUseCase(GetUserRepository(), GetMessageQueueRepository())
}
