package app

import (
	"os"

	"github.com/dionisiopro/dobet_payment/controller"
	"github.com/dionisiopro/dobet_payment/database"
	"github.com/dionisiopro/dobet_payment/event"
	"github.com/dionisiopro/dobet_payment/repository"
	"github.com/dionisiopro/dobet_payment/routes"
	"github.com/dionisiopro/dobet_payment/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer(done <-chan bool) *gin.Engine {
	engine := gin.New()
	conn := RabbitConnection()
	var rabbitMQPublisher = event.NewRabbitMQEventPublisher(conn)
	var collection = database.OpenCollection("payment")
	var repository = repository.NewPaymentMongodbReposiotry(collection)
	var service = service.NewPaymentService(repository, rabbitMQPublisher, nil)
	var controller = controller.NewPaymnetController(service)
	var router = routes.NewPaymentRouter(controller)
	router.SetupPaymentRouter(engine)

	var listenner = event.NewEventListenner()
	var eventSubscriber = event.NewEventSubscriber(conn)
	var betcreatedEvenTListenner = event.NewBetCreatedListenner(eventSubscriber, service)
	var userbetEvenTListenner = event.NewUserBetWinListenner(eventSubscriber, service)
	var userupdatedEvenTListenner = event.NewUserUpdatedListenner(eventSubscriber, service)
	var userdeletedEvenTListenner = event.NewUserDeletedListenner(eventSubscriber, service)
	var usercreatedEvenTListenner = event.NewUserCreatedListenner(eventSubscriber, service)

	listenner.AddListenner(betcreatedEvenTListenner)
	listenner.AddListenner(userbetEvenTListenner)
	listenner.AddListenner(userupdatedEvenTListenner)
	listenner.AddListenner(userdeletedEvenTListenner)
	listenner.AddListenner(usercreatedEvenTListenner)

	go listenner.Listenning(done)
	
	return engine
}

func RabbitConnection() *amqp.Connection {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	url := os.Getenv("AMQP_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	return conn
}
