package app

import (
	"os"

	"github.com/dionisiopro/dobet-bet/controller"
	"github.com/dionisiopro/dobet-bet/database"
	"github.com/dionisiopro/dobet-bet/event"
	"github.com/dionisiopro/dobet-bet/event/listenner"
	"github.com/dionisiopro/dobet-bet/repository"
	"github.com/dionisiopro/dobet-bet/router"
	"github.com/dionisiopro/dobet-bet/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer(done <-chan bool) *gin.Engine {
	//server
	var rabbitConn = RabbitConn()
	var collection = database.OpenCollection("bet")
	var repository = repository.NewBetReposiotry(collection)
	var subscriber = event.NewRabbitMQEventSubscriber(rabbitConn)
	var publisher = event.NewRabbitMQEventPublisher(rabbitConn)
	var service = service.NewBetService(repository, publisher)
	var betConfirmpaymentListenner = listenner.NewconfirmPaymentEventListenner(service, subscriber)
	var betConfirmMatchListenner = listenner.NewConfirmMatchEventListenner(service, subscriber)
	var betMatchResultListenner = listenner.NewmatchResultEventListenner(service, subscriber)
	eventListennermanager := event.NewEventListennersManager(*publisher)
	eventListennermanager.AddListenner(betConfirmMatchListenner)
	eventListennermanager.AddListenner(betMatchResultListenner)
	eventListennermanager.AddListenner(betConfirmpaymentListenner)
	go eventListennermanager.Listenning(done)

	//web
	engine := gin.New()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var controller = controller.NewBetController(service)
	var ginrouter = router.NewBetRouter(controller)
	ginrouter.SetupBetRoutes(engine)
	return engine
}

func RabbitConn() *amqp.Connection {
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	
	adress := os.Getenv("RABBITMQ_URL_HOST")
	conn, err := amqp.Dial(adress)
	if err != nil {
		panic(err)
	}
	return conn

}
