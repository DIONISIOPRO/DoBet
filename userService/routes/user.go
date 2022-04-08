package routes

import (
	"github.com/gin-gonic/gin"
)

type (
	userRoute struct {
		controller UserController
		middleware UsermiddleWare
	}
	UserController interface {
		GetUsers() gin.HandlerFunc
		GetUserById() gin.HandlerFunc
		DeleteUser() gin.HandlerFunc
		UpdateUser() gin.HandlerFunc
	}
	UsermiddleWare interface {
		Authenticated() gin.HandlerFunc
		IsAdmin() gin.HandlerFunc
		IsOwner() gin.HandlerFunc
	}
)

func NewUserRouter(controller UserController, middleware UsermiddleWare) *userRoute {
	return &userRoute{
		controller: controller,
		middleware: middleware,
	}
}

func (router userRoute) SetupUserRouter(app *gin.Engine) *gin.Engine {
	middleware := router.middleware
	app.GET("/api/v1/user", middleware.Authenticated(), middleware.IsAdmin(), router.controller.GetUsers())
	app.GET("/api/v1/user/:id", middleware.Authenticated(), middleware.IsOwner(), router.controller.GetUserById())
	app.POST("/api/v1/user/delete/:id", middleware.Authenticated(), middleware.IsAdmin(), router.controller.DeleteUser())
	app.PUT("/api/v1/user/update/:id", middleware.Authenticated(), middleware.IsOwner(), router.controller.UpdateUser())
	return app
}
