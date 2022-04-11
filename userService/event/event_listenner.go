package event

import (
	"encoding/json"
	"github/namuethopro/dobet-user/domain"
	"log"
	"time"

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
	subscriber EventSubscreber
	processor  EventProcessor
}

func NewRabbitMQEventListenner(processor EventProcessor, subscriber EventSubscreber) EventListenner {
	return EventListenner{
		processor:  processor,
		subscriber: subscriber,
	}
}

func (listenner EventListenner) ListenningToqueues(done <-chan bool) {
	for _, queue := range domain.QueuesToListenning {
		topic, err := listenner.subscriber.SubscribeToQueue(queue)
		if err != nil {
			log.Print(err)
		}
		switch queue {
		case domain.USERDEPOSIT, domain.USERWIN:
			go processMessage(topic, listenner.processor.AddBalance, done)
		case domain.USERWITHDRAW, domain.USERBET:
			go processMessage(topic, listenner.processor.SubtractBalance, done)
		case domain.USERREQUESTBET, domain.USERREQUESTWITHDRAW:
			go processMessage(topic, listenner.processor.CheckMoney, done)
		case domain.USERCREATED:
			go processMessage(topic, printInconsole, done)
		}

	}
}

func printInconsole(data []byte) error {
	usercreated := domain.UserCreatedEvent{}
	err := json.Unmarshal(data, &usercreated)
	if err != nil {
		log.Print("error unmarshaling")
	}
	time.Sleep(time.Second * 2)
	log.Printf("the phone number of user is: %s", usercreated.User.Phone_number)
	return nil
}

func processMessage(queue <-chan amqp.Delivery, processor func([]byte) error, done <-chan bool) {
	goroutinesCountChann := make(chan int, 5)
	for q := range queue {
		select {
		case <-done:
			return
		default:
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
}
