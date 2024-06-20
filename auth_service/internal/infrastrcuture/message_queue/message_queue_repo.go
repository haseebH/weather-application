package message_queue

import (
	"encoding/json"
	"fmt"
	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"github.com/streadway/amqp"
	"log"
)

type rabbitMQPublisher struct {
	connection   *amqp.Connection
	exchangeName string
	routingKey   string
}

func NewRabbitMQPublisher(uri, exchangeName, routingKey string) repository.MessageQueue {
	fmt.Println(uri)
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	err = declareExchange(ch, exchangeName, "direct", true)
	if err != nil {
		log.Fatal(err)
	}
	return &rabbitMQPublisher{
		connection:   conn,
		exchangeName: exchangeName,
		routingKey:   routingKey,
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
			return nil
		}
		return err
	}

	fmt.Printf("Exchange '%s' created.\n", exchangeName)
	return nil
}

// Publish
// todo: wait till message is consumed
func (p *rabbitMQPublisher) Publish(user *repository.User) error {
	ch, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = ch.Publish(
		p.exchangeName, // exchange
		p.routingKey,   // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}

func (p *rabbitMQPublisher) Close() error {
	return p.connection.Close()
}
