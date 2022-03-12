package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/service"
)

type GetBetUrl struct {
	Page    int `uri:"page"`
	PerPage int `uri:"perpage"`
}

type GetBetsByUser struct {
	Page    int    `uri:"page"`
	PerPage int    `uri:"perpage"`
	UserId  string `uri:"id"`
}

func GetBets() gin.HandlerFunc {
	return func(c *gin.Context) {
		parameters := GetBetUrl{}
		c.ShouldBindUri(&parameters)
		bets, err := service.BetService.Bets(int64(parameters.Page), int64(parameters.PerPage))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, bets)
	}
}

func GetBetsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		parameters := GetBetsByUser{}
		err := c.ShouldBindUri(&parameters)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Pleasse Provide page, perpage and userid to access this endpoint"})
		}
		page := parameters.Page
		id := parameters.UserId
		perpage := parameters.PerPage
		bets, err := service.BetService.BetByUser(id, int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, bets)
	}
}

func CreateBet() gin.HandlerFunc {
	return func(c *gin.Context) {
		bet := models.Bet{}
		if err := c.Bind(&bet); err != nil {
			msg := "Error while creating the bet, please provide a valid bet"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		}
		err := service.BetService.CreateBet(bet)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, bet)
	}
}
