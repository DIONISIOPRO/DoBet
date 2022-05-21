package router

import "github.com/gin-gonic/gin"

type BetController interface{
	GetBetsByUserId() gin.HandlerFunc 
	GetBets() gin.HandlerFunc
	CreateBet() gin.HandlerFunc
}
type betRoutes struct {
	controller BetController
}

func NewBetRouter(controller BetController) *betRoutes {
	return &betRoutes{
		controller: controller,
	}
}

func (route *betRoutes) SetupBetRoutes(app *gin.Engine){
	app.GET("/api/v1/bet",route.controller.GetBets())
	app.GET("/api/v1/bet/:id",route.controller.GetBetsByUserId())
	app.POST("/api/v1/bet", route.controller.CreateBet())
}