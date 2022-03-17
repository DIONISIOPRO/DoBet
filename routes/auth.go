package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type authRoutes struct {
	controller controller.AuthController
}

func NewAuthRoute(controller controller.AuthController) *authRoutes {
	return &authRoutes{
		controller: controller,
	}
}
func (route *authRoutes) SetupAuthRoutes(app gin.Engine) gin.Engine {
	app.POST("/login/", route.controller.LogIn())
	app.POST("/logout/", middleware.Authenticated(),route.controller.Logout())
	app.POST("/refresh/", route.controller.Refresh())
	app.POST("/signup/", route.controller.SignUp())
	return app
}
