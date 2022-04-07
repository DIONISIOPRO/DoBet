package message

import (
	"github.com/streadway/amqp"
)

type Event interface {
	ToByteArray() ([]byte, error)
}

type RMQEventManager struct {
	ListenningChannel *amqp.Channel
	PublishingChannel *amqp.Channel
}

func NewRMQEventManager(ListenningChannel *amqp.Channel, PublishingChannel *amqp.Channel) *RMQEventManager {
	return &RMQEventManager{
		ListenningChannel: ListenningChannel,
		PublishingChannel: PublishingChannel,
	}
}

func (manager *RMQEventManager) CreateQueues(queues []string) (err error) {
	for _, queueName := range queues {
		_, err = manager.PublishingChannel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)

	}
	return err
}

func (manager *RMQEventManager) SubscribeToQueue(name string) (<-chan amqp.Delivery, error) {
	return manager.ListenningChannel.Consume(name, "", false, false, false, false, nil)
}

func (manager *RMQEventManager) Publish(name string, event interface{}) error {
	data, err := event.(Event).ToByteArray()
	if err != nil {
		return err
	}
	err = manager.PublishingChannel.Publish(
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

func (manager *RMQEventManager) ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte) error) {
	go func() {
		processMessage(queue, f)
	}()
}

func processMessage(queue <-chan amqp.Delivery, f func([]byte) error) {
	for q := range queue {
		go func(delivery amqp.Delivery) {
			data := delivery.Body
			err := f(data)
			if err != nil {
				delivery.Ack(false)
			}
		}(q)
	}
}
