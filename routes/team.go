package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type teamRoute struct {
	controller controller.TeamController
}

func NewTeamRouter(controller controller.TeamController) *teamRoute {
	return &teamRoute{
		controller: controller,
	}
}

func (route *teamRoute) SetupTeamRouter(app *gin.Engine) *gin.Engine{
    app.GET("/api/v1/team/",middleware.Authenticated(), route.controller.GetTeams())
	app.GET("/api/v1/team/:country", middleware.Authenticated(),route.controller.GetTeamsByCountry())
	return app
}
