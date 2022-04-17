package app

import (
	"github.com/streadway/amqp"
)

func RabbitChannel() (*amqp.Channel, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		return &amqp.Channel{}, &amqp.Channel{}, err
	}
	listenning, err := conn.Channel()
	if err != nil {
		return &amqp.Channel{}, &amqp.Channel{}, err
	}
	publishing, err := conn.Channel()
	if err != nil {
		return &amqp.Channel{}, &amqp.Channel{}, err
	}
	return listenning, publishing, nil
}
