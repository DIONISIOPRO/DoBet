package app

import (
	"github.com/namuethopro/dobet-user/controller"
	"github.com/namuethopro/dobet-user/database"
	"github.com/namuethopro/dobet-user/routes"
	"github.com/namuethopro/dobet-user/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer(done <-chan bool) *gin.Engine {
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
	var service = service.NewService(collection, conn)
	var controller = controller.NewController(service)
	var router = routes.NewRouter(controller)
	engine = router.SetupUserRouter(engine)
	go func() {
		service.StartListenningEvents(done)
	}()
	return engine
}

func RabbitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		panic(err)
	}
	return conn
}
