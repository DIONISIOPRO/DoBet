package controller

import (
	"github/namuethopro/dobet-auth/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	InvalidCredencialsErr = "invalid credential"
	TokenInvalidErr       = "token invalid"
)

type AuthController struct {
	authService AuthService
}
type AuthService interface {
	Login(domain.LoginDetails) (token, refreshToken string, err error)
	Logout(token string) error
	RefreshToken(token string) (acessToken, refreshToken string, err error)
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := controller.authService
		userlogin := domain.LoginDetails{}
		err := c.BindJSON(&userlogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user invalid"})
		}
		token, refreshtoken, err := authService.Login(userlogin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user invalid"})
			c.Abort()
			return
		}
		acessToken := c.Request.Header.Get("token")
		if acessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
			c.Abort()
			return
		}
		err = controller.authService.Logout(acessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token":        acessToken,
			"refreshToken": refreshToken})
	}
}
