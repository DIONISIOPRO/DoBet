package event

import (
	"github/namuethopro/dobet-user/domain"
	"log"

	"github.com/streadway/amqp"
)

type EventProcessor interface {
	AddBalance(data []byte) error
	SubtractBalance(data []byte) error
	CheckMoney(data []byte) error
}
type EventSubscreber interface {
	SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
}
type EventListenner struct {
	subscriber        EventSubscreber
	processor         EventProcessor
	ListenningChannel *amqp.Channel
}

func NewRabbitMQEventListenner(processor EventProcessor, subscriber EventSubscreber, ListenningChannel *amqp.Channel) *EventListenner {
	return &EventListenner{
		processor:         processor,
		subscriber:        subscriber,
		ListenningChannel: ListenningChannel,
	}
}

func (listenner *EventListenner) ListenningToqueues() {
	for _, queue := range domain.QueuesToListenning {
		topic, err := listenner.subscriber.SubscribeToQueue(queue)
		if err != nil {
			log.Print(err.Error())
		}
		switch queue {
		case domain.USERDEPOSIT, domain.USERWIN:
			go processMessage(topic, listenner.processor.AddBalance)
		case domain.USERWITHDRAW, domain.USERBET:
			go processMessage(topic, listenner.processor.SubtractBalance)
		case domain.USERREQUESTBET, domain.USERREQUESTWITHDRAW:
			go processMessage(topic, listenner.processor.CheckMoney)
		}

	}
}

func processMessage(queue <-chan amqp.Delivery, processor func([]byte) error) {
	goroutinesCountChann := make(chan int, 10)
	for q := range queue {
		goroutinesCountChann <- 1
		go func(delivery amqp.Delivery) {
			data := delivery.Body
			err := processor(data)
			if err != nil {
				delivery.Ack(false)
			}
			<-goroutinesCountChann
		}(q)
	}
}
