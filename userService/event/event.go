package event

import (
	"log"

	"github.com/streadway/amqp"
)

type Event interface {
	ToByteArray(obj interface{}) []byte
}

type RMQEventManager struct {
	ListenningChannel *amqp.Channel
	PublishingChannel *amqp.Channel
}

func NewRMQEventManager(ListenningChannel *amqp.Channel,PublishingChannel *amqp.Channel) *RMQEventManager {
	return &RMQEventManager{
		ListenningChannel: ListenningChannel,
		PublishingChannel: PublishingChannel,
	}
}

func (manager *RMQEventManager) CreateQueues(queues []string) (err error) {
	for _, queueName := range queues{
		_, err = manager.PublishingChannel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		
	}
	return  err
}

func (manager *RMQEventManager) SubscribeToQueue(name string) (<-chan amqp.Delivery, error) {
	return manager.ListenningChannel.Consume(name, "", true, false, false, false, nil)
}

func (manager *RMQEventManager) Publish(name string, body []byte) error {
	err := manager.PublishingChannel.Publish(
		"", name, false, false,
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: 2,
			Body:         body,
		},
	)
	if err != nil{
		return err
	}
	log.Print("event published")
	return nil

}

func (manager *RMQEventManager) ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte)) {
	go func() {
		for q := range queue {
			data := q.Body
			f(data)
		}
	}()
}


