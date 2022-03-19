package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"

)

type userRoute struct {
	controller controller.UserController
}

func NewUserRouter(controller controller.UserController) *userRoute {
	return &userRoute{
		controller: controller,
	}
}


func (router *userRoute) SetupUserRouter(app *gin.Engine) *gin.Engine{
	app.GET("/api/v1/user", middleware.Authenticated(), middleware.OnlyForAdmin(), router.controller.GetUsers())
	app.GET("/api/v1/user/:id",middleware.Authenticated(), middleware.IfIdParamIsOwner(), router.controller.GetUserById())
	return app
}