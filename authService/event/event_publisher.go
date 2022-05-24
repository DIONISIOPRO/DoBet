package event

import (
	"errors"
	"log"

	"github.com/dionisiopro/dobet-auth/domain"

	"github.com/streadway/amqp"
)

type EventPublisher struct {
	PublishingChannel *amqp.Channel
}

func NewRabbitMQEventPublisher(conn *amqp.Connection) EventPublisher {
	channel, err := conn.Channel()
	if err != nil{
		log.Print("can`t create a channel to publish")
		panic(err)
	}
	return EventPublisher{
		PublishingChannel: channel,
	}
}

func (publisher EventPublisher) Publish(name string, event domain.Event) error {
	if name == "" || event == nil {
		return errors.New("invalid parameters")
	}
	data, err := event.ToByteArray()
	if err != nil {
		return err
	}
	publisher.PublishingChannel.QueueDeclare(name, false, false, false, false, nil)
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
