package routes

import (
	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/middleware"
	"github.com/gin-gonic/gin"
)

type userRoute struct {
	controller controller.UserController
	middleware middleware.JWTMiddleWare
}

func NewUserRouter(controller controller.UserController, middleware middleware.JWTMiddleWare) *userRoute {
	return &userRoute{
		controller: controller,
		middleware: middleware,
	}
}


func (router userRoute) SetupUserRouter(app *gin.Engine) *gin.Engine{
	middleware := router.middleware
	app.GET("/api/v1/user", middleware.Authenticated(),  middleware.IsAdmin(), router.controller.GetUsers())
	app.GET("/api/v1/user/:id",middleware.Authenticated(), middleware.IsOwner(), router.controller.GetUserById())
	app.PUT("/api/v1/user/delete/:id", middleware.Authenticated(), middleware.IsAdmin(), router.controller.DeleteUser())
	app.POST("/api/v1/user/update/:id", middleware.Authenticated(), middleware.IsOwner(), router.controller.UpdateUser())
	return app
}