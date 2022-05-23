package controller

import (
	"github.com/dionisiopro/dobet-auth/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	InvalidCredencialsErr = "invalid credential"
	TokenInvalidErr       = "token invalid"
)

type LoginSucess struct {
	Token       string `json:"token"`
	RefresToken string `json:"refresh_toten"`
}
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

// Login godoc
// @Summary Login in the system
// @Description this route allows you to login in the dobet server
// @Accept  json
// @Produce  json
// @Success 200 {object} LoginSucess	"tokens"
// @Failure 500 {string} string :"error"
// @Param       user    body     domain.LoginDetails     true  "credentials"
// @Router /login [post]
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
		credencials := LoginSucess{Token: token, RefresToken: refreshtoken}
		c.JSON(http.StatusOK, credencials)

	}
}

// LogOut godoc
// @Summary LogOut in the system
// @Description this route allows you to LogOut in the dobet server
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.LogoutDetails	"credentials"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Param        user   body   domain.LogoutDetails    true  "credentials"
// @Router /logout [post]
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

// RefreshToken godoc
// @Summary LogOut in the system
// @Description this route allows you to LogOut in the dobet server
// @Accept  json
// @Produce  json
// @Success 200 {object} LoginSucess	"credentials"
// @Failure 400 {string} string :"error"
// @Failure 500 {string} string :"error"
// @Param       user body     domain.LogoutDetails    true  "credentials"
// @Router /logout [post]
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
		credencials := LoginSucess{Token: acessToken, RefresToken: refreshToken}
		c.JSON(http.StatusOK, credencials)
	}
}
