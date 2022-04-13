package controller

import (
	"encoding/json"
	"github/namuethopro/dobet-user/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		User_id      string    `json:"user_id"`
		First_name   string    `json:"first_name"`
		Last_name    string    `json:"last_name"`
		Phone_number string    `json:"phone_number"`
		Created_at   time.Time `json:"created_at"`
		IsAdmin      bool      `json:"is_admin"`
	}
	ResponseList []Response

	LoginRequest struct {
		Phone_number string `json:"phone_number" validate:"required"`
		Password     string `json:"password" validate:"required"`
	}
	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	Controller struct {
		service Service
	}

	Service interface {
		GetUsers(page, perpage int64) ([]domain.User, error)
		GetUserById(userId string) (domain.User, error)
		GetUserByPhone(phone string) (domain.User, error)
		DeleteUser(userid string) error
		UpdateUser(userid string, user domain.User) error
	}
)

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}


// GetUsers godoc
// @Summary Get list of users
// @Description Get users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Param user body models.AddUser true "Add user"
// @Failure 500 {object} models.Error
// @Failure 400 {object} models.Error
// @Success 200 {object} models.Message
// @Router /users [post]
func (controller *Controller) GetUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}
	perpage, err := strconv.Atoi(c.Query("perpage"))
	if err != nil {
		perpage = 0
	}
	users, err := controller.service.GetUsers(int64(page), int64(perpage))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	usersResponse := ResponseList{}
	usersResponse = usersResponse.FromUsers(users)
	c.JSON(http.StatusOK, usersResponse)
}

func (controller *Controller) GetUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}
	user, err := controller.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userResponse := Response{}
	userResponse.FromUser(user)
	c.JSON(http.StatusOK, userResponse)
}
func (controller *Controller) GetUserByPhone(c *gin.Context) {
	phone := c.Param("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}
	user, err := controller.service.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userResponse := Response{}
	userResponse.FromUser(user)
	c.JSON(http.StatusOK, userResponse)
}

func (controller *Controller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}
	err := controller.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sucess": "User Deleted"})
}

func (controller *Controller) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}
	var user = domain.User{}
	body := c.Request.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user details"})
		return
	}
	err = controller.service.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sucess": "User updated"})
}

func (userResponse *Response) FromUser(user domain.User) *Response {
	response := userToResponse(user)
	return &response
}

func (usersResponse ResponseList) FromUsers(users []domain.User) ResponseList {
	for _, user := range users {
		usersResponse = append(usersResponse, userToResponse(user))
	}
	return usersResponse
}

func userToResponse(user domain.User) Response {
	userResponse := Response{}
	userResponse.Created_at = user.Created_at
	userResponse.First_name = user.First_name
	userResponse.IsAdmin = user.IsAdmin
	userResponse.Last_name = user.Last_name
	userResponse.Phone_number = user.Phone_number
	userResponse.User_id = user.User_id
	return userResponse
}
