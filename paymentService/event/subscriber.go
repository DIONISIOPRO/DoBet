package event

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQEventSubscriber struct{
	Conn      *amqp.Connection
}

func (s RabbitMQEventSubscriber) SubscribeToQueue(name string) <-chan amqp.Delivery{
	channel, err := s.Conn.Channel()
		if err != nil{
			log.Print("error creating channel")
		}
		_, err = channel.QueueDeclare(name, true, false,false, false,nil)
		if err != nil{
			log.Print("error declaring queue")
		}
		queue, err := channel.Consume(name, "", true, false, false, false, nil)
		if err != nil {
			log.Print("error subscribing in to queue: ", err.Error())
			panic(err)
		}
		return queue
}