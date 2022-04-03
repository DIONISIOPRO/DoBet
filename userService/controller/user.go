package controller

import (
	"github/namuethopro/dobet-user/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		userService UserService
	}
	UserService interface {
		GetUsers(page, perpage int64) (domain.UsersResponse, error)
		GetUserById(userId string) (domain.UserResponse, error)
		GetUserByPhone(phone string) (domain.UserResponse, error)
		DeleteUser(userid string) error
		UpdateUser(userid string, user domain.User) error
	}
)

func NewUserController(userService UserService) *UserController {
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
		users, err := controller.userService.GetUsers(int64(page), int64(perpage))
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
		user := domain.User{}
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user details"})
			return
		}
		id := c.Param("id")
		err = controller.userService.UpdateUser(id, user)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, gin.H{"sucess": "User updated"})
	}
}
