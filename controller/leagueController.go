package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)

type GetLeagueUrl struct {
	Page    int `uri:"page"`
	PerPage int `uri:"perpage"`
}

type GetLeagueId struct {
	Page    int    `uri:"page"`
	PerPage int    `uri:"perpage"`
	Country  string `uri:"country"`
}

func GetLeagues() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetLeagueUrl{}
		c.ShouldBindUri(&param)

		leagues, err := service.LeagueService.Leagues(int64(param.Page), int64(param.PerPage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, leagues)
	}
}

func GetLeaguesByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetLeagueId{}
	if err :=	c.ShouldBindUri(&param); err != nil{
		msg := "Pleasse Provide a valid country name of the league"
		c.JSON(http.StatusBadRequest, gin.H{"Error":msg})
	}
		leagues, err := service.LeagueService.GetLeaguesByCountry(param.Country,int64(param.Page), int64(param.PerPage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, leagues)
	}
}
