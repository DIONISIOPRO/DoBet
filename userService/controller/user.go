package controller

import (
	"github/namuethopro/dobet-user/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	UserResponse struct {
		User_id      string    `json:"user_id"`
		First_name   string    `json:"first_name"`
		Last_name    string    `json:"last_name"`
		Phone_number string    `json:"phone_number"`
		Created_at   time.Time `json:"created_at"`
		IsAdmin      bool      `json:"is_admin"`
	}
	UsersResponse []UserResponse

	UserLoginRequest struct {
		Phone_number string `json:"phone_number" validate:"required"`
		Password     string `json:"password" validate:"required"`
	}
	UserLoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	UserController struct {
		userService UserService
	}

	UserService interface {
		GetUsers(page, perpage int64) ([]domain.User, error)
		GetUserById(userId string) (domain.User, error)
		GetUserByPhone(phone string) (domain.User, error)
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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usersResponse := UsersResponse{}
		usersResponse = usersResponse.FromUsers(users)
		c.JSON(http.StatusOK, usersResponse)
	}
}

func (controller *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := controller.userService.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userResponse := UserResponse{}
		userResponse.FromUser(user)
		c.JSON(http.StatusOK, userResponse)
	}
}

func (controller *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := controller.userService.DeleteUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sucess": "User updated"})
	}
}

func (userResponse *UserResponse) FromUser(user domain.User) *UserResponse {
	response := userToResponse(user)
	return &response
}

func (usersResponse UsersResponse) FromUsers(users []domain.User) UsersResponse {
	for _, user := range users {
		usersResponse = append(usersResponse, userToResponse(user))
	}
	return usersResponse
}

func userToResponse(user domain.User) UserResponse {
	userResponse := UserResponse{}
	userResponse.Created_at = user.Created_at
	userResponse.First_name = user.First_name
	userResponse.IsAdmin = user.IsAdmin
	userResponse.Last_name = user.Last_name
	userResponse.Phone_number = user.Phone_number
	userResponse.User_id = user.User_id
	return userResponse
}
