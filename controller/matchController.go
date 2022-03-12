package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)

type GetMatchUrl struct {
	Page    int `uri:"page"`
	PerPage int `uri:"perpage"`
}

type GetMatchLeagueId struct {
	Page    int    `uri:"page"`
	PerPage int    `uri:"perpage"`
	league  string `uri:"league"`
}
type GetMatchLeagueIddDay struct {
	Page    int    `uri:"page"`
	PerPage int    `uri:"perpage"`
	league  string `uri:"league"`
	day     int    `uri:"day"`
}

func GetMatches() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetMatchUrl{}
		c.ShouldBindUri(&param)

		matches, err := service.MatchService.Matches(int64(param.Page), int64(param.PerPage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, matches)
	}
}

func GetMatchesByLeague() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetMatchLeagueId{}
		if err := c.ShouldBindUri(&param); err != nil {
			msg := "Please provide a valid league id"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
		}

		matches, err := service.MatchService.MatchesByLeagueId(param.league, int64(param.Page), int64(param.PerPage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, matches)
	}
}

func GetMatchesLeagueAndday() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := GetMatchLeagueIddDay{}
		if err := c.ShouldBindUri(&param); err != nil {
			msg := "Please provide a valid league id and day"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
		}

		matches, err := service.MatchService.MatchesByLeagueIdDay(param.league, int64(param.day),int64(param.Page), int64(param.PerPage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, matches)
	}
}
