package app

import (
	"os"

	"github.com/dionisiopro/dobet-auth/controller"
	"github.com/dionisiopro/dobet-auth/database"
	"github.com/dionisiopro/dobet-auth/event"
	"github.com/dionisiopro/dobet-auth/repository"
	"github.com/dionisiopro/dobet-auth/routes"
	"github.com/dionisiopro/dobet-auth/service"
	"github.com/dionisiopro/dobet-auth/token"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer(done chan bool) *gin.Engine {
	engine := gin.New()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var conn = RabbitConn()
	var SECRETE_KEY = os.Getenv("JWT_SECRETE_KEY")
	var PrivateKey = []byte(SECRETE_KEY)
	var collection = database.OpenCollection("auth")
	var authRepo = repository.NewAuthRepository(collection)
	var tokenManager = token.NewTokenManager(PrivateKey)
	var eventPublisher = event.NewRabbitMQEventPublisher(conn)
	var service = service.NewAuthService(authRepo,tokenManager,eventPublisher)

	var controller = controller.NewAuthController(service)
	var router = routes.NewAuthRouter(controller)
	router.SetupAuthRoutes(engine)

	var listenner = event.NewAuthEventListenner()
	var subscriber = event.NewRabbitMQEventSubscriber(conn)
	var usercreatedEventListenner = event.NewUserCreatedEventListenner(subscriber,service)
	var userupdatedEventListenner = event.NewUseruUpdatedEventListenner(subscriber, service)
	var userdeletedEventListenner = event.NewUseruDeletedEventListenner(subscriber,service)
     
	listenner.AddListenner(usercreatedEventListenner)
	listenner.AddListenner(userdeletedEventListenner)
	listenner.AddListenner(userupdatedEventListenner)

	go listenner.Listenning(done)
	return engine
}

func RabbitConn() *amqp.Connection {
	adress := "amqp://localhost:5672"
	conn, err := amqp.Dial(adress)
	if err != nil {
		panic(err)
	}
	return conn

}
