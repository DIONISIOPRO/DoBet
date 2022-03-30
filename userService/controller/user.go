package controller

import (
	"net/http"
	"strconv"
	"github/namuethopro/dobet-user/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) GetUsers() gin.HandlerFunc {
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
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, users)
	}
}

func (controller *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := controller.userService.GetUserById(id)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, user)
	}
}

func (controller *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := controller.userService.DeleteUser(id)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, gin.H{"sucess": "User Deleted"})
	}
}

func (controller *UserController) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := controller.userService.UpdateUser(id)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, gin.H{"sucess": "User updated"})
	}
}
