package message_queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/haseebh/weatherapp_temperature/internal/entities/repository"
	"github.com/haseebh/weatherapp_temperature/internal/usecases"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type rabbitMQConsumer struct {
	connection    *amqp.Connection
	exchangeName  string
	routingKey    string
	queueName     string
	temperatureUC usecases.TemperatureUseCase
}

func NewRabbitMQConsumer(temUC usecases.TemperatureUseCase, uri, exchangeName, routingKey, queueName string) repository.MessageQueue {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	err = declareExchange(ch, exchangeName, "direct", true)
	if err != nil {
		log.Fatalf("Failed to declare or check the exchange '%s': %v", exchangeName, err)
	}
	err = declareQueue(ch, queueName, true)
	if err != nil {
		log.Fatalf("Failed to declare or check the exchange '%s': %v", exchangeName, err)
	}
	// Bind the queue to the exchange with the routing key
	err = ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		log.Fatalf("Failed to bind queue '%s' to exchange '%s' with routing key '%s': %v", queueName, exchangeName, routingKey, err)
	}
	return &rabbitMQConsumer{
		connection:    conn,
		exchangeName:  exchangeName,
		routingKey:    routingKey,
		queueName:     queueName,
		temperatureUC: temUC,
	}
}
func declareExchange(ch *amqp.Channel, exchangeName, exchangeType string, durable bool) error {
	// Attempt to declare the exchange
	err := ch.ExchangeDeclare(exchangeName, exchangeType, durable, false, false, false, nil)
	if err != nil {
		// Check if the error is due to the exchange already existing
		if amqpErr, ok := err.(*amqp.Error); ok && amqpErr.Code == amqp.PreconditionFailed {
			// Exchange already exists, so ignore the error
			fmt.Printf("Exchange '%s' already exists.\n", exchangeName)

		} else {
			// Any other error, return it
			return err
		}
	}
	return nil
}

func declareQueue(ch *amqp.Channel, queueName string, durable bool) error {
	// Attempt to declare the queue
	_, err := ch.QueueDeclare(queueName, durable, false, false, false, nil)
	if err != nil {
		// Check if the error is due to the queue already existing
		if amqpErr, ok := err.(*amqp.Error); ok && amqpErr.Code == amqp.PreconditionFailed {
			// Queue already exists, so ignore the error
			fmt.Printf("Queue '%s' already exists.\n", queueName)
			return nil
		}
		// Any other error, return it
		return err
	}

	// Queue created successfully
	fmt.Printf("Queue '%s' created.\n", queueName)
	return nil
}

// Consume
// todo: wait till message is consumed
func (p *rabbitMQConsumer) Consume() error {
	ch, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		p.queueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			var user repository.User
			err := json.Unmarshal(d.Body, &user)
			if err != nil {
				log.Printf("Error decoding message: %s", err)
				continue
			}
			endTime := time.Now().UTC()
			sTime := time.Date(endTime.Year()-3, endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.UTC)
			err = p.temperatureUC.FetchAndStoreTemperature(context.Background(), user.Location, sTime, endTime)
			if err != nil {
				log.Printf("Error fetching weather: %s", err)
				continue
			}
		}
	}()
	return nil
}

func (p *rabbitMQConsumer) Close() error {
	return p.connection.Close()
}
