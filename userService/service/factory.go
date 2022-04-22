package service

import (
	"github.com/namuethopro/dobet-user/event"
	"github.com/namuethopro/dobet-user/repository"
	"sync"

	"github.com/streadway/amqp"
)

func NewService(collection repository.MongoDriverUserCollection, conn *amqp.Connection) userService {
	var repository = repository.NewUserRepository(collection)
	var moneyReserver = newMoneyReserver(&sync.Mutex{})
	var eventpublisher = event.NewRabbitMQEventPublisher(conn)
	var eventsubscriber = event.NewRabbitMQEventSubscriber(conn)
	var eventProcessor = event.NewIncomingEventProcessor(&sync.Mutex{}, repository, eventpublisher, moneyReserver)
	var eventlistenner = event.NewRabbitMQEventListenner(eventProcessor, eventsubscriber)
	return newUserService(repository, eventpublisher, eventlistenner, eventProcessor, &sync.Mutex{})
}
