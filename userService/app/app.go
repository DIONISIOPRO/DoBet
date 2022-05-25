package app

import (
	"os"

	"github.com/dionisiopro/dobet-user/controller"
	"github.com/dionisiopro/dobet-user/database"
	"github.com/dionisiopro/dobet-user/event"
	"github.com/dionisiopro/dobet-user/repository"
	"github.com/dionisiopro/dobet-user/routes"
	"github.com/dionisiopro/dobet-user/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer() *gin.Engine {
	engine := gin.New()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	conn := RabbitConnection()
	if err != nil {
		panic(err)
	}
	var collection = database.OpenCollection("users")
	var userRepo = repository.NewUserRepository(collection)
	var publisher = event.NewRabbitMQEventPublisher(conn)
	var service = service.NewUserService(userRepo,publisher)
	var controller = controller.NewController(service)
	var router = routes.NewRouter(controller)
	engine = router.SetupUserRouter(engine)
	return engine
}

func RabbitConnection() *amqp.Connection {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	adress := os.Getenv("RABBITMQ_URL_HOST")
	conn, err := amqp.Dial(adress)
	if err != nil {
		panic(err)
	}
	return conn
}
