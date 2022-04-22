package service

import (
	"github.com/namuethopro/dobet-auth/token"
	"github.com/namuethopro/dobet-auth/event"
	"github.com/namuethopro/dobet-auth/repository"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewService(privatekey []byte, collection *mongo.Collection, listenningChannel, publishingChannel *amqp.Channel) authService {
	repo := repository.NewAuthRepository(collection)
	jwtManager := token.NewTokenManager(privatekey)
	eventmanger := event.NewEventManger(publishingChannel, listenningChannel, collection)
	service := newAuthService(repo, &eventmanger, jwtManager, PasswordHandler{})
	return *service
}
