package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type leagueRoute struct {
	controller controller.LeagueController
}

func NewLeagueRouter(controller controller.LeagueController) *leagueRoute{
	return &leagueRoute{
		controller: controller,
	}
}

func (route *leagueRoute) SetupLeagueRouter(app *gin.Engine) *gin.Engine{
	app.GET("/api/v1/league/",middleware.Authenticated(), route.controller.GetLeagues())
	app.GET("/api/v1/league/:country", middleware.Authenticated(),route.controller.GetLeaguesByCountry())
	return app
}