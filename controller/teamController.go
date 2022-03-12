package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)


type GetTeamUrl struct {
	Page    int `uri:"page"`
	PerPage int `uri:"perpage"`
}

type GetTeamLeagueId struct {
	Page    int    `uri:"page"`
	PerPage int    `uri:"perpage"`
	Country  string `uri:"country"`
}
func GetTeams() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetTeamUrl{}
		c.ShouldBindUri(&param)
		teams, err := service.TeamService.Teams(int64(param.Page),int64(param.PerPage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, teams)
	}
}

func GetTeamsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetTeamLeagueId{}
		if err := c.ShouldBindUri(&param); err != nil {
			msg := "Please provide a valid league id"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
		}

		teams, err := service.TeamService.TeamsByCountry(param.Country, int64(param.Page), int64(param.PerPage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, teams)
	}
}
