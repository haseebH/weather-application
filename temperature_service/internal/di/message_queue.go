package di

import (
	"github.com/haseebh/weatherapp_temperature/internal/entities/repository"
	"github.com/haseebh/weatherapp_temperature/internal/infrastrcuture/message_queue"
	"github.com/haseebh/weatherapp_temperature/pkg/config"
)

func GetMessageQueueRepository() repository.MessageQueue {
	cfg := config.LoadConfig()
	return message_queue.NewRabbitMQConsumer(
		GetTemperatureUseCase(),
		cfg.MessageQueue.RabbitMQURI,
		cfg.MessageQueue.ExchangeName,
		cfg.MessageQueue.RoutingKey,
		cfg.MessageQueue.QueueName)
}
