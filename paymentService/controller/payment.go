package controller

import (
	"net/http"

	"github.com/dionisiopro/dobet_payment/domain"
	"github.com/gin-gonic/gin"
)

type SucessResponse struct {
	Sucess string `json:"sucess"`
}
type ErrorResponse struct {
	ErrorMsg string `json:"error"`
}
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

// Deposit godoc
// @Summary deposit your money
// @Description this route allows you to deposit your money
// @Accept  json
// @Produce  json
// @Success 200 {object} SucessResponse	"sucess"
// @Failure 500 {object} ErrorResponse "error"
// @Param    deposit body domain.Deposit  true "deposit"
// @Router /deposit [get]
func (controller *PaymentController) Deposit(c *gin.Context) {
	deposit := domain.Deposit{}
	err := c.BindJSON(&deposit)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{ErrorMsg: "Invalid amount or user Id"})
		return
	}
	if err = controller.service.Deposit(deposit); err != nil {
		c.JSON(http.StatusInternalServerError,  ErrorResponse{ErrorMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SucessResponse{Sucess: "Your deposit was succefly"})
}

// WithDraw godoc
// @Summary WithDraw your money
// @Description this route allows you to WithDraw your money
// @Accept  json
// @Produce  json
// @Success 200 {object} SucessResponse	"sucess"
// @Failure 500 {object} ErrorResponse "error"
// @Param    WithDraw body domain.WithDraw  true "WithDraw"
// @Router /withdraw [get]
func (controller *PaymentController) WithDraw(c *gin.Context) {
	withdraw := domain.WithDraw{}
	err := c.BindJSON(&withdraw)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{ErrorMsg: "Invalid amount or user Id"})
		return
	}
	if err = controller.service.WithDraw(withdraw); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{ErrorMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK,  SucessResponse{Sucess: "Your withdraw was succefly"})
}
