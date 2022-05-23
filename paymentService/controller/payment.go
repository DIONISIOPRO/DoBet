package controller

import (
	"net/http"

	"github.com/dionisiopro/dobet_payment/domain"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Deposit(domain.Deposit) error
	WithDraw(domain.WithDraw) error
}
type PaymentController struct {
	service Service
}

func NewPaymnetController(service Service) *PaymentController {
	return &PaymentController{
		service: service,
	}
}

func (controller *PaymentController) Deposit(c *gin.Context) {
	deposit := domain.Deposit{}
	err := c.BindJSON(&deposit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount or user Id"})
		return
	}
	if err = controller.service.Deposit(deposit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sucess": "Your deposit was succefly"})
}

func (controller *PaymentController) WithDraw(c *gin.Context) {
	withdraw := domain.WithDraw{}
	err := c.BindJSON(&withdraw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount or user Id"})
		return
	}
	if err = controller.service.WithDraw(withdraw); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sucess": "Your withdraw was succefly"})
}
