package app

import (
	"os"
	
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
	var SECRETE_KEY = os.Getenv("JWT_SECRETE_KEY")
	var PrivateKey = []byte(SECRETE_KEY)
	var collection = database.OpenCollection("auth")
	var service = service.NewService(PrivateKey, collection, RabbitConn())
	var controller = controller.NewAuthController(&service)
	var router = routes.NewAuthRouter(controller)
	engine = router.SetupAuthRoutes(engine)
	go func() {
		service.StartListenningToEvents(done)
	}()
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
