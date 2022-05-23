package service

import (
	"github.com/dionisiopro/dobet-user/event"
	"github.com/dionisiopro/dobet-user/repository"
	"sync"

	"github.com/streadway/amqp"
)

func NewService(collection repository.MongoDriverUserCollection, conn *amqp.Connection) userService {
	var repository = repository.NewUserRepository(collection)
	var eventpublisher = event.NewRabbitMQEventPublisher(conn)
	var eventsubscriber = event.NewRabbitMQEventSubscriber(conn)
	var eventProcessor = event.NewIncomingEventProcessor(repository)
	var eventlistenner = event.NewRabbitMQEventListenner(eventProcessor, eventsubscriber)
	return newUserService(repository, eventpublisher, eventlistenner, eventProcessor, &sync.Mutex{})
}
