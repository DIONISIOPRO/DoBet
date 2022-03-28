package routes

import (
	"github/namuethopro/dobet-user/controller"
	"github/namuethopro/dobet-user/middleware"

	"github.com/gin-gonic/gin"
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