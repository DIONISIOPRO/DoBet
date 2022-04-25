package event

import (
	"log"

	"github.com/dionisiopro/dobet_payment/domain"
)

type EventProcessor interface {
	CreateUser([]byte) error
	UpdateUser([]byte) error
	DeleteUser([]byte) error
	Pay([]byte) error
	Deposit([]byte) error
}
type RabbitMQListenner struct {
	conn      *amqp.Connection
	processor EventProcessor
}

func (l *RabbitMQListenner) Listenning(done <-chan bool) {
	for _, event := range domain.EventsToListenning {
		channel := l.conn.Channel()
		queue, err := channel.Consume(event, "", true, false, false, false, nil)
		if err != nil {
			log.Printf("error subscribing in to queue: ", err.Error())
		}
		switch event {
		case domain.USERBETCREATED:
			go processMessage(queue, l.processor.Pay, done)
		case domain.USERCREATED:
			go processMessage(queue, l.processor.CreateUser, done)
		case domain.USERUPDATED:
			go processMessage(queue, l.processor.UpdateUser, done)
		case domain.USERBETWIN:
			go processMessage(queue, l.processor.Deposit, done)
		case domain.USERDELETED:
			go processMessage(queue, l.processor.DeleteUser, done)
		default:
			continue
		}

	}
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
