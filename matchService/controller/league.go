package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)
type LeagueController struct{
	leagueService service.LeagueService
}

func NewLeagueController(leagueService service.LeagueService) *LeagueController{
	return &LeagueController{
		leagueService: leagueService,
	}
}
func (controller *LeagueController)GetLeagues() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}

		leagues, err := controller.leagueService.Leagues(int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()

		}
		c.JSON(http.StatusOK, leagues)
	}
}

func (controller *LeagueController)GetLeaguesByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}
		country := c.Param("country")
		
		leagues, err := controller.leagueService.GetLeaguesByCountry(country, int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
		}
		c.JSON(http.StatusOK, leagues)
	}
}
