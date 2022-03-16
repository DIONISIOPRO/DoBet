package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)

type matchController struct {
	matchService service.MatchService
}

func NewMatchController(matchService service.MatchService) *matchController {
	return &matchController{
		matchService: matchService,
	}
}

func (controller *matchController)GetMatchesLeagueAndday() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}
		league := c.Param("league")
		day, err := strconv.Atoi(c.Param("perpage"))
		if err != nil {
			day = 0
		}
		matches, err := controller.matchService.MatchesByLeagueIdDay(league, int64(day), int64(page), int64(perpage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
			return
		}
		c.JSON(http.StatusOK, matches)
	}
}
