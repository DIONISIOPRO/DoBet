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
type LoginError struct{
	Msg string `json:"error"`
}
type AuthController struct {
	authService AuthService
}
type AuthService interface {
	Login(user domain.LoginDetails) (token, refreshToken string, err error)
	Logout(token string) error
	RefreshToken(token string) (acessToken, refreshToken string, err error)
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login godoc
// @Summary Login in the DoBet
// @Description this route allows you to login in the dobet
// @Accept  json
// @Produce  json
// @Success 200 {object} LoginSucess "This document contain your tokens"
// @Failure 500 {object} LoginError "this document contain the error occured"
// @Param       user    body     domain.LoginDetails     true  "give your login credencials"
// @Router /login [post]
func (controller *AuthController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := controller.authService
		userlogin := domain.LoginDetails{}
		err := c.BindJSON(&userlogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, LoginError{Msg:"user invalid"})
		}
		token, refreshtoken, err := authService.Login(userlogin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, LoginError{Msg: err.Error()})
			return
		}
		credencials := LoginSucess{Token: token, RefresToken: refreshtoken}
		c.JSON(http.StatusOK, credencials)

	}
}

// LogOut godoc
// @Summary Logout in the system
// @Description this route allows you to Logout in the dobet
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.LogoutDetails	"this document contain your login credentials"
// @Failure 400 {object} LoginError "this document contain the error occured"
// @Failure 500 {object} LoginError "this document contain the error occured"
// @Param        credentials   body   domain.LogoutDetails    true  "give your login credentials"
// @Router /logout [post]
func (controller *AuthController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutUser := domain.LogoutDetails{}
		err := c.BindJSON(&logoutUser)
		if err != nil {
			c.JSON(http.StatusBadRequest,LoginError{Msg: "user invalid"})
			c.Abort()
			return
		}
		acessToken := c.Request.Header.Get("token")
		if acessToken == "" {
			c.JSON(http.StatusUnauthorized, LoginError{Msg: "token is empty"})
			c.Abort()
			return
		}
		err = controller.authService.Logout(acessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, LoginError{Msg: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, logoutUser)
	}
}

// RefreshToken godoc
// @Summary Get you refresh tokens
// @Description this route allows you to request new tokens if your token ispirex
// @Accept  json
// @Produce  json
// @Success 200 {object} LoginSucess	"this document contain your login credentials"
// @Failure 401 {object} LoginError "this document contain the error occured"
// @Failure 500 {object} LoginError "this document contain the error occured"
// @Param       token  header     string   true  "give your expired token"
// @Router /refresh [post]
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
