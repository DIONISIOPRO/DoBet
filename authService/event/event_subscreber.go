package event

import (
	"log"

	"github.com/streadway/amqp"
)

type EventSubscriber struct {
	ListenningChannel *amqp.Channel
}

func NewRabbitMQEventSubscriber(ListenningChannel *amqp.Channel) EventSubscriber {
	return EventSubscriber{
		ListenningChannel: ListenningChannel,
	}
}

func (manager EventSubscriber) SubscribeToQueue(name string) (<-chan amqp.Delivery, error) {
	manager.ListenningChannel.QueueDeclare(name, false, false, false, false, nil)
	queue , err :=  manager.ListenningChannel.Consume(name, "", true, false, false, false, nil)
	if err != nil{
		log.Printf("error subscribing to queue: %v", err)
	}
	return queue, nil
}
