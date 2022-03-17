package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/service"
)

type BetController struct{
	betService service.BetService
}

func NewBetController(betService service.BetService) *BetController{
	return &BetController{
		betService: betService,
	}
}

func (controller *BetController) GetBets() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}
		bets, err := controller.betService.Bets(int64(page),int64(perpage))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bets)
	}
}

func(controller *BetController) GetBetsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}
		id := c.Param("id")
		bets, err := controller.betService.BetByUser(id, int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bets)
	}
}

func (controller *BetController)CreateBet() gin.HandlerFunc {
	return func(c *gin.Context) {
		bet := models.Bet{}
		if err := c.BindJSON(&bet); err != nil {
			msg := "Error while creating the bet, please provide a valid bet"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		err := controller.betService.CreateBet(bet)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bet)
	}
}
