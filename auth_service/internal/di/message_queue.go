package di

import (
	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"github.com/haseebh/weatherapp_auth/internal/infrastrcuture/message_queue"
	"github.com/haseebh/weatherapp_auth/pkg/config"
)

func GetMessageQueueRepository() repository.MessageQueue {
	cfg := config.LoadConfig()
	return message_queue.NewRabbitMQPublisher(cfg.MessageQueue.RabbitMQURI, cfg.MessageQueue.ExchangeName, cfg.MessageQueue.RoutingKey)
}
