package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/namuethopro/dobet-user/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

type (
	userRoute struct {
		controller Controller
	}
	Controller interface {
		GetUsers(c *gin.Context)
		GetUserById(c *gin.Context)
		GetUserByPhone(c *gin.Context)
		DeleteUser(c *gin.Context)
		UpdateUser(c *gin.Context)
	}
)

func NewRouter(controller Controller) *userRoute {
	return &userRoute{
		controller: controller,
	}
}

func (router userRoute) SetupUserRouter(app *gin.Engine) *gin.Engine {
	app.GET("/api/v1/user",router.controller.GetUsers)
	app.GET("/api/v1/user/:id", router.controller.GetUserById)
	app.GET("/api/v1/user/phone/:phone", router.controller.GetUserByPhone)
	app.POST("/api/v1/user/delete/:id",  router.controller.DeleteUser)
	app.PUT("/api/v1/user/update/:id",  router.controller.UpdateUser)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
