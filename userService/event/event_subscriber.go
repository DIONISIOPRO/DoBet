package event

import (
	"log"

	"github.com/streadway/amqp"
)

type EventSubscriber struct {
	Conn *amqp.Connection
}

func NewRabbitMQEventSubscriber(conn *amqp.Connection) *EventSubscriber {
	return &EventSubscriber{
		Conn: conn,
	}
}

func (manager EventSubscriber) SubscribeToQueue(name string) (<-chan amqp.Delivery, error) {
	channel, err := manager.Conn.Channel()
	if err != nil{
		log.Printf("error creating channel: %v", err)
	}
	queue , err :=  channel.Consume(name, "", true, false, false, false, nil)
	if err != nil{
		log.Printf("error subscribing to queue: %v", err)
	}
	return queue, nil
}
