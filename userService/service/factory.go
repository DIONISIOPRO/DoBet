package service

import (
	"fmt"
	"github/namuethopro/dobet-user/event"
	"github/namuethopro/dobet-user/repository"
	"sync"

	"github.com/streadway/amqp"
)

type outroListenner struct{}

func (outroListenner) ListenningToqueues() {
	fmt.Print("funcionando")
}

var este = outroListenner{}

func NewService(collection repository.MongoDriverUserCollection, PublishingChannel, ListenningChannel *amqp.Channel) userService {
	var repository = repository.NewUserRepository(collection)
	var moneyReserver = newMoneyReserver(&sync.Mutex{})
	var eventpublisher = event.NewRabbitMQEventPublisher(PublishingChannel)
	var eventsubscriber = event.NewRabbitMQEventSubscriber(ListenningChannel)
	var eventProcessor = event.NewIncomingEventProcessor(&sync.Mutex{}, repository, eventpublisher, moneyReserver)
	var eventlistenner = event.NewRabbitMQEventListenner(eventProcessor, eventsubscriber)
	return newUserService(repository, eventpublisher, eventlistenner, eventProcessor, &sync.Mutex{})
}
