package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)
type teamController struct{
	teamService service.TeamService
}

func NewTeamRepository(teamService service.TeamService) *teamController{
	return &teamController{
		teamService: teamService,
	}
}
func (controller *teamController)GetTeams() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil{
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil{
			perpage = 0
		}
		teams, err := controller.teamService.Teams(int64(page),int64(perpage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, teams)
	}
}

func(controller *teamController) GetTeamsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil{
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil{
			perpage = 0
		}
		country := c.Param("country")
		teams, err := controller.teamService.TeamsByCountry(country, int64(page), int64(perpage))
		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		c.JSON(http.StatusOK, teams)
	}
}
