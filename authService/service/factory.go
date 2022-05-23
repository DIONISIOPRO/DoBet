package service

import (
	"github.com/dionisiopro/dobet-auth/event"
	"github.com/dionisiopro/dobet-auth/repository"
	"github.com/dionisiopro/dobet-auth/token"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewService(privatekey []byte, collection *mongo.Collection, conn *amqp.Connection) authService {
	repo := repository.NewAuthRepository(collection)
	tokenManger := token.NewTokenManager(privatekey)
	eventmanger := event.NewEventManger(conn, collection)
	service := newAuthService(repo, &eventmanger, tokenManger, PasswordHandler{})
	return *service
}
