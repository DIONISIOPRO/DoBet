package controller

import (
	"encoding/json"
	"github.com/dionisiopro/dobet-user/domain"
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
	UserResponseList []UserResponse

	LoginRequest struct {
		Phone_number string `json:"phone_number" validate:"required"`
		Password     string `json:"password" validate:"required"`
	}
	ResponseError struct {
		Msg        string `json:"error"`
	}
	SuccesResponse struct{
		Msg string `json:"sucess"`
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
// @Summary Get a list of users <<only for admin>>
// @Description if you are admin, and want get all users use this route to  get a list of users given a number of page and limit of users per page
// @Accept  json
// @Produce  json
// @Success 200 {object} UserResponseList	"list of users"
// @Failure 500 {object} ResponseError "error"
// @Param        page         query     int     false  "give the ppage number"       minimum(1)
// @Param        perpage         query     int     false  "give the limit per page"       minimum(9)    maximum(20)
// @Router /users [get]
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
		c.JSON(http.StatusInternalServerError, ResponseError{Msg: err.Error()})
		return
	}
	usersResponse := UserResponseList{}
	usersResponse = usersResponse.FromUsers(users)
	c.JSON(http.StatusOK, usersResponse)
}

// GetUserById godoc
// @Summary get a user by id
// @Description get user by ID
// @Accept  json
// @Produce  json
// @Param   id      path   int     true  "user id"
// @Success 200 {object} UserResponse	"user"
// @Failure 400 {object} ResponseError "msg error"
// @Failure 500 {object} ResponseError "msg of error"
// @Router /users/{id} [get]
func (controller *Controller) GetUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Msg: "invalid id param"})
		return
	}
	user, err := controller.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Msg: err.Error()})
		return
	}
	userResponse := UserResponse{}
	userResponse.FromUser(user)
	c.JSON(http.StatusOK, userResponse)
}

// GetUserByPhone godoc
// @Summary get a user by phone
// @Description get user by phone
// @Accept  json
// @Produce  json
// @Param   phone      path   int     true  "user phone"
// @Success 200 {object} UserResponse	"ok"
// @Failure 400 {object} ResponseError "error msg"
// @Failure 500 {object} ResponseError "error msg"
// @Router /users/{phone} [get]
func (controller *Controller) GetUserByPhone(c *gin.Context) {
	phone := c.Param("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Msg: "invalid phone param"})
		return
	}
	user, err := controller.service.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Msg: err.Error()})
		return
	}
	userResponse := UserResponse{}
	userResponse.FromUser(user)
	c.JSON(http.StatusOK, userResponse)
}

// DeleteUser
// @Summary delete a user by id
// @Description delete user by ID
// @Accept  json
// @Produce  json
// @Param   id      path   int     true  "user id"
// @Success 200 {object} SuccesResponse	"sucess message"
// @Failure 400 {object} ResponseError "error message"
// @Failure 500 {object} ResponseError "error message"
// @Router /users/delete/{id} [delete]
func (controller *Controller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}
	err := controller.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccesResponse{Msg: "User Deleted"})
}

// UpdateUser godoc
// @Summary update a user by id
// @Description update user by ID
// @Accept  json
// @Produce  json
// @Param   id      path   int     true  "user id"
// @Param   user      body domain.User true  "Some id"
// @Success 200 {object} SuccesResponse	"sucess message"
// @Failure 400 {object} ResponseError "error message"
// @Failure 500 {object} ResponseError "error message"
// @Router /users/update/{id} [put]
func (controller *Controller) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Msg: "invalid id param"})
		return
	}
	var user = domain.User{}
	body := c.Request.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Msg: "invalid user details"})
		return
	}
	err = controller.service.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccesResponse{Msg: "User updated"})
}

func (userResponse *UserResponse) FromUser(user domain.User) {
	response := userToResponse(user)
	userResponse = &response
}

func (usersResponse UserResponseList) FromUsers(users []domain.User) UserResponseList {
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
