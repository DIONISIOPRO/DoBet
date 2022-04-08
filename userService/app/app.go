package app

import (
	"os"
	"sync"

	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/database"
	"github/namuethopro/dobet-user/message"
	"github/namuethopro/dobet-user/middleware"
	"github/namuethopro/dobet-user/repository"
	"github/namuethopro/dobet-user/routes"
	"github/namuethopro/dobet-user/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Application struct {
	Engine *gin.Engine
}

func (application *Application) Run(host string) {
	application.Engine.Run(host)
}

func (application *Application) Setup(engine *gin.Engine) *Application {
	application.Engine = engine
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	lock := &sync.Mutex{}
	listenningchannel, publishingChannel, err := RabbitChannel()
	if err != nil {
		panic(err)
	}

	var SECRETE_KEY = os.Getenv("JWT_SECRETE_KEY")
	var PrivateKey = []byte(SECRETE_KEY)
	var rabbitEventManager = message.NewRMQEventManager(listenningchannel, publishingChannel)
	var logoutManager = service.NewLogInStateManager()
	var jwtmiddleware = middleware.NewjwtMiddleWare(logoutManager, PrivateKey)
	var userCollection = database.OpenCollection("users")
	var userRepository = repository.NewUserRepository(userCollection)
	var moneyreserver = service.NewMoneyReserver(lock)

	var incomingEventHandler = service.NewIncomingEventHandler(lock, userRepository, rabbitEventManager, &moneyreserver)
	var userService = service.NewUserService(userRepository, rabbitEventManager, incomingEventHandler, lock)
	var userController = controller.NewUserController(userService)
	var userRouter = routes.NewUserRouter(userController, jwtmiddleware)
	application.Engine = userRouter.SetupUserRouter(application.Engine)
	return application
}

func RabbitChannel() (*amqp.Channel, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		return &amqp.Channel{}, &amqp.Channel{}, err
	}
	listenning, err := conn.Channel()
	if err != nil {
		return &amqp.Channel{}, &amqp.Channel{}, err
	}
	publishing, err := conn.Channel()
	if err != nil {
		return &amqp.Channel{}, &amqp.Channel{}, err
	}
	return listenning, publishing, nil
}
