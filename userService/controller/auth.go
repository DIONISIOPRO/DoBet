package controller

import (
	"github/namuethopro/dobet-user/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	InvalidCredencialsErr = "invalid credential"
	TokenInvalidErr       = "token invalid"
)

type (
	AuthController struct {
		authService AuthService
	}
	AuthService interface {
		Login(domain.LoginDetails) (token, refreshToken string, err error)
		Logout(token string) error
		SignUp(userRequest domain.UserSignUpRequest) (string, error)
		RefreshToken(token string) (acessToken, refreshToken string, err error)
	}
)

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}
func (controller *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := domain.UserSignUpRequest{}
		err := c.BindJSON(&user)
		checkBadRequestErr(c, err, InvalidCredencialsErr)
		userid, err := controller.authService.SignUp(user)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusCreated, gin.H{
			"id":    userid,
			"phone": user.Phone_number,
		})
	}
}

func (controller *AuthController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := controller.authService
		userlogin := domain.LoginDetails{}
		err := c.BindJSON(&userlogin)
		checkBadRequestErr(c, err, InvalidCredencialsErr)
		token, refreshtoken, err := authService.Login(userlogin)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusAccepted, gin.H{
			"token":        token,
			"refreshToken": refreshtoken,
		})

	}
}

func (controller *AuthController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutUser := domain.LogoutDetails{}
		err := c.BindJSON(&logoutUser)
		checkBadRequestErr(c, err, InvalidCredencialsErr)
		acessToken := c.Request.Header.Get("token")
		if acessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token iss empty"})
			return
		}
		err = controller.authService.Logout(acessToken)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, logoutUser)
	}
}

func (controller *AuthController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := controller.authService
		acessToken := c.Request.Header.Get("token")
		if acessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
			return
		}
		acessToken, refreshToken, err := authService.RefreshToken(acessToken)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusOK, gin.H{
			"token":        acessToken,
			"refreshToken": refreshToken})
	}
}

func checkInternalServerErr(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func checkBadRequestErr(c *gin.Context, err error, msg string) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
		return
	}
}
