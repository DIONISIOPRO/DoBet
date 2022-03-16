package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/service"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}

		users, err := controller.userService.Users(int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})

		}
		c.JSON(http.StatusOK, users)
	}
}

func (controller *userController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := controller.userService.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
	}
}
