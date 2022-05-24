package event

import "github.com/streadway/amqp"

type RabbiMQEventPublisher struct {
	channel *amqp.Channel
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
