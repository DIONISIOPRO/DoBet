package router

import (
	_ "github.com/dionisiopro/dobet-bet/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

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
	app.GET("/api/v1/swagger/*any",  ginSwagger.WrapHandler(swaggerFiles.Handler))
}