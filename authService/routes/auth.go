package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/namuethopro/dobet-auth/docs"
)

type (
	authRoutes struct {
		controller authController
	}
	authController interface {
		LogIn() gin.HandlerFunc
		Logout() gin.HandlerFunc
		Refresh() gin.HandlerFunc
	}
)

func NewAuthRouter(controller authController) *authRoutes {
	return &authRoutes{
		controller: controller,
	}
}
func (route *authRoutes) SetupAuthRoutes(app *gin.Engine) *gin.Engine {
	app.POST("/api/v1/login", route.controller.LogIn())
	app.POST("/api/v1/logout", route.controller.Logout())
	app.POST("/api/v1/refresh", route.controller.Refresh())
	app.GET("/api/v1/swagger/*any",  ginSwagger.WrapHandler(swaggerFiles.Handler))
	return app
}
