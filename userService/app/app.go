package app

import (
	"os"

	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/database"
	"github/namuethopro/dobet-user/middleware"
	"github/namuethopro/dobet-user/routes"
	"github/namuethopro/dobet-user/service"

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
	listenningchannel, publishingChannel, err := RabbitChannels()
	if err != nil {
		panic(err)
	}
	var SECRETE_KEY = os.Getenv("JWT_SECRETE_KEY")
	var PrivateKey = []byte(SECRETE_KEY)
	var logoutManager = service.NewLogInStateManager()
	var middleware = middleware.NewjwtMiddleWare(logoutManager, PrivateKey)
	var collection = database.OpenCollection("users")
	var service = service.NewService(collection, publishingChannel, listenningchannel)
	var controller = controller.NewController(service)
	var router = routes.NewRouter(controller, middleware)
	engine = router.SetupUserRouter(engine)
	go func() {
		service.StartListenningEvents()
	}()
	return engine
}

func RabbitChannels() (*amqp.Channel, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		panic(err)
	}
	listenning, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	publishing, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return listenning, publishing, nil
}
