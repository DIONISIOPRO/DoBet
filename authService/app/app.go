package app

import (
	"os"
	"time"

	"github.com/namuethopro/dobet-auth/controller"
	"github.com/namuethopro/dobet-auth/database"
	"github.com/namuethopro/dobet-auth/routes"
	"github.com/namuethopro/dobet-auth/service"

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
	listenningchannel, publishingChannel, err := RabbitChannel()
	if err != nil {
		panic(err)
	}
	var SECRETE_KEY = os.Getenv("JWT_SECRETE_KEY")
	var PrivateKey = []byte(SECRETE_KEY)
	var collection = database.OpenCollection("auth")
	var service = service.NewService(PrivateKey, collection, listenningchannel, publishingChannel)
	var controller = controller.NewAuthController(&service)
	var router = routes.NewAuthRouter(controller)
	engine = router.SetupAuthRoutes(engine)
	go func() {
		service.StartEventHandler(done)
	}()
	return engine
}

func RabbitChannel() (*amqp.Channel, *amqp.Channel, error) {
	adress := "amqp://localhost:5672"
	listenningConn, err := amqp.Dial(adress)
	if err != nil {
		panic(err)
	}
	time.Sleep(3*time.Second)
	time.Sleep(3*time.Second)
	listenning, err := listenningConn.Channel()
	if err != nil {
		panic(err)
	}
	pubConn, err := amqp.Dial(adress)
	if err != nil {
		panic(err)
	}
	time.Sleep(3* time.Second)

	publishing, err := pubConn.Channel()
	if err != nil {
		panic(err)
	}
	return listenning, publishing, nil
}
