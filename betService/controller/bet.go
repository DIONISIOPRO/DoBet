package controller

import (
	"net/http"
	"strconv"

	"github.com/dionisiopro/dobet-bet/domain/bet"
	"github.com/gin-gonic/gin"
)

type BetService interface {
	CreateBet(bet *bet.BetBase) (string, error)
	BetByUser(user_id string, page, perpage int64) ([]bet.BetBase, error)
	Bets(page, perpage int64) ([]bet.BetBase, error)
}

type BetController struct {
	betService BetService
}

func NewBetController(betService BetService) *BetController {
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
		bets, err := controller.betService.Bets(int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"bets": bets})
	}
}

func (controller *BetController) GetBetsByUserId() gin.HandlerFunc {
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
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"bets": bets})
	}
}

func (controller *BetController) CreateBet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bet bet.BetBase

		if err := c.BindJSON(&bet); err != nil {
			msg := "Error while creating the bet, please provide a valid bet"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			c.Abort()
			return
		}
		bet_id, err := controller.betService.CreateBet(&bet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"bet_id": bet_id})
		c.Abort()
	}
}
