package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/dionisiopro/dobet-auth/docs"
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