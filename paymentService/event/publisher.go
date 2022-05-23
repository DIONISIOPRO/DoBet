package event

import "github.com/streadway/amqp"

type RabbiMQEventPublisher struct {
	channel *amqp.Channel
}

func NewRabbitMQEventPublisher(conn *amqp.Connection) *RabbiMQEventPublisher {
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return &RabbiMQEventPublisher{
		channel: channel,
	}
}

func (publisher RabbiMQEventPublisher) Publish(name string, data []byte) error {
	publisher.channel.QueueDeclare(name, false, false, false, false, nil)
	err := publisher.channel.Publish(
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
