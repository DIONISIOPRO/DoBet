package routes

import (
	"github.com/gin-gonic/gin"
)

type (
	authRoutes struct {
		controller authController
		middleware authmiddleWare
	}
	authmiddleWare interface {
		Authenticated() gin.HandlerFunc
	}
	authController interface {
		SignUp() gin.HandlerFunc
		LogIn() gin.HandlerFunc
		Logout() gin.HandlerFunc
		Refresh() gin.HandlerFunc
	}
)

func NewAuthRouter(controller authController, middleware authmiddleWare) *authRoutes {
	return &authRoutes{
		controller: controller,
		middleware: middleware,
	}
}
func (route *authRoutes) SetupAuthRoutes(app *gin.Engine) *gin.Engine {
	middleware := route.middleware
	app.POST("/api/v1/login", route.controller.LogIn())
	app.POST("/api/v1/logout", middleware.Authenticated(), route.controller.Logout())
	app.POST("/api/v1/refresh", route.controller.Refresh())
	app.POST("/api/v1/signup", route.controller.SignUp())
	return app
}
