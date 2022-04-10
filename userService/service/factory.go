package service

import (
	"github/namuethopro/dobet-user/event"
	"github/namuethopro/dobet-user/repository"
	"sync"

	"github.com/streadway/amqp"
)

func NewService(collection repository.MongoDriverUserCollection, PublishingChannel, ListenningChannel *amqp.Channel) *userService {
	repository := repository.NewUserRepository(collection)
	moneyReserver := newMoneyReserver(&sync.Mutex{})
	eventpublisher := event.NewRabbitMQEventPublisher(PublishingChannel)
	eventsubscriber := event.NewRabbitMQEventSubscriber(ListenningChannel)
	eventProcessor := event.NewIncomingEventProcessor(&sync.Mutex{}, repository, eventpublisher, moneyReserver)
	eventlistenner := event.NewRabbitMQEventListenner(eventProcessor, eventsubscriber)
	return newUserService(repository, eventpublisher, eventlistenner, eventProcessor, &sync.Mutex{})
}
