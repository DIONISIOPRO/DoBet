package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type betRoutes struct {
	controller controller.BetController
}

func NewBetRouter(controller controller.BetController) *betRoutes {
	return &betRoutes{
		controller: controller,
	}
}

func (route *betRoutes) SetupBetRoutes(app *gin.Engine) *gin.Engine{
	app.GET("/api/v1/bet", middleware.Authenticated(),route.controller.GetBets())
	app.GET("/api/v1/bet/:id", middleware.Authenticated(),route.controller.GetBetsByUserId())
	app.POST("/api/v1/bet", middleware.Authenticated(),route.controller.CreateBet())
	return app
}