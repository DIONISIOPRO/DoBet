package app

import (
	"os"

	"github.com/dionisiopro/dobet_payment/controller"
	"github.com/dionisiopro/dobet_payment/database"
	"github.com/dionisiopro/dobet_payment/repository"
	"github.com/dionisiopro/dobet_payment/routes"
	"github.com/dionisiopro/dobet_payment/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func CreateGinServer(done <-chan bool) *gin.Engine {
	engine := gin.New()
	//conn := RabbitConnection()
	var collection = database.OpenCollection("payment")
	var repository = repository.NewPaymentMongodbReposiotry(collection)
	var service = service.NewPaymentService(repository, nil)
	var controller = controller.NewPaymnetController(service)
	var router = routes.NewPaymentRouter(controller)
	engine = router.SetupPaymentRouter(engine)
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
