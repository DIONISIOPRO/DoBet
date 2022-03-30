package routes

import (
	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/middleware"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	controller controller.AuthController
	middleware middleware.JWTMiddleWare
}

func NewAuthRouter(controller controller.AuthController, 	middleware middleware.JWTMiddleWare) *authRoutes {
	return &authRoutes{
		controller: controller,
		middleware: middleware,
	}
}
func (route *authRoutes) SetupAuthRoutes(app *gin.Engine) *gin.Engine {
	middleware := route.middleware
	app.POST("/api/v1/login", route.controller.LogIn())
	app.POST("/api/v1/logout", middleware.Authenticated(),route.controller.Logout())
	app.POST("/api/v1/refresh", route.controller.Refresh())
	app.POST("/api/v1/signup", route.controller.SignUp())
	return app
}
