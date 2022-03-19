package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type authRoutes struct {
	controller controller.AuthController
}

func NewAuthRouter(controller controller.AuthController) *authRoutes {
	return &authRoutes{
		controller: controller,
	}
}
func (route *authRoutes) SetupAuthRoutes(app *gin.Engine) *gin.Engine {
	app.POST("/api/v1/login", route.controller.LogIn())
	app.POST("/api/v1/logout", middleware.Authenticated(),route.controller.Logout())
	app.POST("/api/v1/refresh",middleware.Authenticated(), route.controller.Refreshh())
	app.POST("/api/v1/signup", route.controller.SignUp())
	return app
}
