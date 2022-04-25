package app

import (
	"os"

	"github.com/namuethopro/dobet_payment/controller"
	"github.com/namuethopro/dobet_payment/database"
	"github.com/namuethopro/dobet_payment/routes"
	"github.com/namuethopro/dobet_payment/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer(done <-chan bool) *gin.Engine {
	engine := gin.New()
	conn := RabbitConnection()
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
