package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/service"
)

type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymnetController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

func (controller *PaymentController) Deposit() gin.HandlerFunc {
	return func(c *gin.Context) {
		deposit := models.Deposit{}
		err := c.BindJSON(&deposit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount or user Id"})
			return
		}
		if err = controller.paymentService.Deposit(deposit.Amount, deposit.UserId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sucess": "Your deposit was succefly"})
	}
}

func (controller *PaymentController) WithDraw() gin.HandlerFunc {
	return func(c *gin.Context) {
		withdraw := models.WithDraw{}
		err := c.BindJSON(&withdraw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount or user Id"})
			return
		}
		if err = controller.paymentService.Deposit(withdraw.Amount, withdraw.UserId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sucess": "Your withdraw was succefly"})
	}
}
