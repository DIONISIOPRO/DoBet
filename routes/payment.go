package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type paymentRoute struct {
	controller controller.PaymentController
}

func NewPaymentRouter(controller controller.PaymentController) *paymentRoute {
	return &paymentRoute{
		controller: controller,
	}
}

func (route *paymentRoute) SetupPaymentRouter(app *gin.Engine) *gin.Engine {
	app.POST("/api/v1/deposit", middleware.Authenticated(),route.controller.Deposit())
	app.POST("/api/v1/withdraw", middleware.Authenticated(),route.controller.WithDraw())
	return app
}
