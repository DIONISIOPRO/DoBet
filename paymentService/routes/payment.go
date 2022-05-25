package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/dionisiopro/dobet_payment/docs"
)

type PaymentController interface {
	Deposit(*gin.Context)
	WithDraw(*gin.Context)
}
type paymentRoute struct {
	controller PaymentController
}

func NewPaymentRouter(controller PaymentController) *paymentRoute {
	return &paymentRoute{
		controller: controller,
	}
}

func (route *paymentRoute) SetupPaymentRouter(app *gin.Engine){
	app.POST("/api/v1/deposit", route.controller.Deposit)
	app.POST("/api/v1/withdraw", route.controller.WithDraw)
	app.GET("/api/v1/swagger/*any",  ginSwagger.WrapHandler(swaggerFiles.Handler))
}
