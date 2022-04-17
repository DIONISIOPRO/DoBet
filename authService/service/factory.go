package service

import (
	"github/namuethopro/dobet-auth/auth"
	"github/namuethopro/dobet-auth/event"
	"github/namuethopro/dobet-auth/repository"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewService(privatekey []byte, collection *mongo.Collection, listenningChannel, publishingChannel *amqp.Channel) authService {
	repo := repository.NewAuthRepository(collection)
	jwtManager := auth.NewJwtManager(privatekey)
	eventmanger := event.NewEventManger(publishingChannel, listenningChannel, collection)
	service := newAuthService(repo, &eventmanger, jwtManager, PasswordHandler{})
	return *service
}
