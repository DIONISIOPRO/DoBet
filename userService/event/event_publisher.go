package event

import (
	"github/namuethopro/dobet-user/domain"

	"github.com/streadway/amqp"
)

type EventPublisher struct {
	PublishingChannel *amqp.Channel
}

func NewEventPiblisher(PublishingChannel *amqp.Channel) *EventPublisher {
	return &EventPublisher{
		PublishingChannel: PublishingChannel,
	}
}

func (publisher *EventPublisher) Publish(name string, event domain.Event) error {
	data, err := event.ToByteArray()
	if err != nil {
		return err
	}
	err = publisher.PublishingChannel.Publish(
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
