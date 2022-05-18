package event

import (
	"errors"

	"github.com/streadway/amqp"
)
type EventPublisher struct {
	PublishingChannel *amqp.Channel
}

func NewRabbitMQEventPublisher(PublishingChannel *amqp.Channel) EventPublisher {
	return EventPublisher{
		PublishingChannel: PublishingChannel,
	}
}

func (publisher EventPublisher) Publish(name string, data []byte) error {
	if name == "" || data == nil {
		return errors.New("invalid parameters")
	}
	publisher.PublishingChannel.QueueDeclare(name, false, false, false, false, nil)
	err := publisher.PublishingChannel.Publish(
		"", name, false, false,
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: 2,
			Body:         data,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
