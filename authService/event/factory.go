package event

import (
	"github/namuethopro/dobet-auth/repository"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEventManger(pchannel, lchannel *amqp.Channel, collection *mongo.Collection) EventManger {
	repo := repository.NewAuthRepository(collection)
	publisher := NewRabbitMQEventPublisher(pchannel)
	processor := NewIncomingEventProcessor(repo)
	creator := NewQueueCreator(pchannel)
	subscriber := NewRabbitMQEventSubscriber(lchannel)
	manager := newEventManager(publisher,processor, creator,subscriber )
	return manager
}