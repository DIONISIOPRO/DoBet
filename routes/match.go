package routes

import (
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/middleware"
)

type matchRoute struct {
	controller controller.MatchController
}

func NewMatchRouter(controller controller.MatchController) *matchRoute {
	return &matchRoute{
		controller: controller,
	}
}

func (route *matchRoute) SetupMatchRouter(app gin.Engine) gin.Engine {
	app.GET("/match/:league/:day",middleware.Authenticated(), route.controller.GetMatchesLeagueAndday())
	return app
}
