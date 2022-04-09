package event

import "github.com/streadway/amqp"

type EventSubscriber struct {
	ListenningChannel *amqp.Channel
}

func NewRabbitMQEventSubscriber(ListenningChannel *amqp.Channel) *EventSubscriber{
	return &EventSubscriber{
		ListenningChannel:ListenningChannel,
	}
}

func (manager *EventSubscriber) SubscribeToQueue(name string) (<-chan amqp.Delivery, error) {
	return manager.ListenningChannel.Consume(name, "", false, false, false, false, nil)
}