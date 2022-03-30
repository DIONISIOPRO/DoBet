package app

import (
	"github/namuethopro/dobet-user/auth"
	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/database"
	"github/namuethopro/dobet-user/middleware"
	"github/namuethopro/dobet-user/repository"
	"github/namuethopro/dobet-user/routes"
	"github/namuethopro/dobet-user/service"

	"github.com/gin-gonic/gin"
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

func (application Application) Setup() *gin.Engine{
	var PrivateKey = []byte("secrete")
	var jwtmanager = auth.NewJwtManager(PrivateKey)
	var logoutManager = service.NewLogoutMangger()
	var jwtmiddleware = middleware.NewJwtMiddleware(jwtmanager, logoutManager)
	var userCollection = database.OpenCollection("users")
	var userRepository = repository.NewUserRepository(userCollection)
	var authRepository = repository.NewAuthRepository(userCollection)
	var userService = service.NewUserService(userRepository)
	var authService = service.NewAuthService(authRepository, logoutManager)
	var userController = controller.NewUserController(userService)
	var authContoller = controller.NewAuthController(authService, jwtmanager)
	var userRouter = routes.NewUserRouter(*userController, jwtmiddleware)
	var authRouter = routes.NewAuthRouter(*authContoller, jwtmiddleware)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app = userRouter.SetupUserRouter(app)
	app = authRouter.SetupAuthRoutes(app)
	
	return app
}

