package routes

import (
	"github.com/gin-gonic/gin"
)

type (
	userRoute struct {
		controller Controller
		middleware MiddleWare
	}
	Controller interface {
		GetUsers(c *gin.Context)
		GetUserById(c *gin.Context)
		DeleteUser(c *gin.Context)
		UpdateUser(c *gin.Context)
	}
	MiddleWare interface {
		Authenticated() gin.HandlerFunc
		IsAdmin() gin.HandlerFunc
		IsOwner() gin.HandlerFunc
	}
)

func NewRouter(controller Controller, middleware MiddleWare) *userRoute {
	return &userRoute{
		controller: controller,
		middleware: middleware,
	}
}

func (router userRoute) SetupUserRouter(app *gin.Engine) *gin.Engine {
	middleware := router.middleware
	app.GET("/api/v1/user", middleware.Authenticated(), middleware.IsAdmin(), router.controller.GetUsers)
	app.GET("/api/v1/user/:id", middleware.Authenticated(), middleware.IsOwner(), router.controller.GetUserById)
	app.GET("/api/v1/user/phone/:phone", middleware.IsAdmin())
	app.POST("/api/v1/user/delete/:id", middleware.Authenticated(), middleware.IsAdmin(), router.controller.DeleteUser)
	app.PUT("/api/v1/user/update/:id", middleware.Authenticated(), middleware.IsOwner(), router.controller.UpdateUser)
	return app
}
