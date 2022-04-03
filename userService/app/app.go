package app

import (
	"os"
	"sync"

	"github/namuethopro/dobet-user/auth"
	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/database"
	"github/namuethopro/dobet-user/event"
	"github/namuethopro/dobet-user/middleware"
	"github/namuethopro/dobet-user/repository"
	"github/namuethopro/dobet-user/routes"
	"github/namuethopro/dobet-user/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Application struct {
	Host string
}

var application = Application{}

func NewApplication(host string) Application {
	application.Host = host
	return application
}
func (application Application) Run() {
	app := application.Setup()
	app.Run(application.Host)
}

func (application Application) Setup() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	lock := &sync.Mutex{}
	listenningchannel, publishingChannel, err := RabbitChannel()
	if err != nil {
		panic(err)
	}
	rabbitEventManager := event.NewRMQEventManager(listenningchannel,publishingChannel)
	var SECRETE_KEY = os.Getenv("JWT_SECRETE_KEY")
	var PrivateKey = []byte(SECRETE_KEY)
	var jwtmanager = auth.NewJwtManager(PrivateKey)
	var logoutManager = service.NewLogoutMangger()
	var jwtmiddleware = middleware.NewjwtMiddleWare(jwtmanager, logoutManager)
	var userCollection = database.OpenCollection("users")
	var userRepository = repository.NewUserRepository(userCollection)
	var authRepository = repository.NewAuthRepository(userCollection)
	var userService = service.NewUserService(userRepository, rabbitEventManager, lock)
	var authService = service.NewAuthService(authRepository, logoutManager, rabbitEventManager,jwtmanager)
	var userController = controller.NewUserController(userService)
	var authContoller = controller.NewAuthController(authService)
	var userRouter = routes.NewUserRouter(userController, jwtmiddleware)
	var authRouter = routes.NewAuthRouter(authContoller, jwtmiddleware)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app = userRouter.SetupUserRouter(app)
	app = authRouter.SetupAuthRoutes(app)

	return app
}

func RabbitChannel() (*amqp.Channel, *amqp.Channel, error) {
	conn, err := amqp.Dial("")
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
