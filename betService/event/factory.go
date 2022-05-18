package event

// import (
// 	"github.com/dionisiopro/dobet-bet/repository"
// 	"github.com/streadway/amqp"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func NewEventManger(conn *amqp.Connection, collection *mongo.Collection) EventManger {
// 	channel, err := conn.Channel()
// 	if err != nil{
// 		panic(err)
// 	}
// 	repo := repository.NewAuthRepository(collection)
// 	publisher := NewRabbitMQEventPublisher(channel)
// 	processor := NewIncomingEventProcessor(repo)
// 	subscriber := NewRabbitMQEventSubscriber(conn)
// 	manager := newEventManager(publisher,processor,subscriber )
// 	return manager
// }