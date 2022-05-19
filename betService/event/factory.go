package event

import (
	"github.com/dionisiopro/dobet-bet/repository"
	"github.com/dionisiopro/dobet-bet/service"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEventManager(conn *amqp.Connection, collection *mongo.Collection) *EventManager {
	channel, err := conn.Channel()
	if err != nil{
		panic(err)
	}
	publisher := NewRabbitMQEventPublisher(channel)
	subscriber := NewRabbitMQEventSubscriber(conn)
	repo := repository.NewBetReposiotry(collection)
	service := service.NewBetService(repo,publisher)
	betConfirmpaymentListenner := ConfirmPaymentEventListenner{
		service: service,
		subscriber: subscriber,
	}
	betConfirmMatchListenner := ConfirmMatchEventListenner{
		service: service,
		subscriber: subscriber,
	}
	betMatchResultListenner := MatchResultEventListenner{
		service: service,
		subscriber: subscriber,
	}
	manager := newEventManager(publisher)
	manager.AddListenner(betConfirmMatchListenner)
	manager.AddListenner(betMatchResultListenner)
	manager.AddListenner(betConfirmpaymentListenner)
	return manager
}