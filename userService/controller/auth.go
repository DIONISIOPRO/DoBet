package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github/namuethopro/dobet-user/domain"
	"github/namuethopro/dobet-user/service"
	"github/namuethopro/dobet-user/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		userlogin := domain.LoginDetails{}
		err := c.BindJSON(&userlogin)
		if err != nil {
			msg := "Please provide an valid login credential"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		user, err := controller.authService.Login(userlogin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		PasswordSameErr := utils.CompareHashedPassword(user.Hashed_password, userlogin.Password)
		if PasswordSameErr != nil {
			msg := "Please provide an valid login credential"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		crsf := utils.GenerateCrsfToken()
		acessToken, err := utils.GenerateToken(crsf, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating acess token"})
			return
		}
		refreshToken, err := utils.GenerateRefreshToken(crsf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cookie := &http.Cookie{
			Name:     "token",
			Value:    acessToken,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)

		ok := controller.authService.UpdateRefreshToken(refreshToken, user.User_id)
		if !ok {
			log.Print("Can not save the refresh token")
		}
		utils.SetCrsfTokenToClient(c.Writer, crsf)
		c.JSON(http.StatusOK, gin.H{"token": acessToken, "refreshtoken": refreshToken})
	}
}

func (controller *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := domain.User{}
		err := c.BindJSON(&user)
		if err != nil {
			msg := "Please povide a valid user"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
			return
		}
		user.Hashed_password, err = utils.HasPassword(user.Password)
		if err != nil {
			msg := "Please povide a valid user"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
			return
		}
		if err := controller.authService.SignUp(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, user.Phone_number)
	}
}

func (controller *AuthController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutUser := domain.LogoutDetails{}
		err := c.BindJSON(&logoutUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		acessToken, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, err := utils.GrabClaimsFromAcessToken(acessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if claims.Phone != logoutUser.Phone {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user credentials"})
			return
		}
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Local().Add(-100 * time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		controller.authService.UpdateRefreshToken("", claims.Uid)
		c.JSON(http.StatusOK, logoutUser)
	}
}

func (controller *AuthController) Refreshh() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("starting")
		token := utils.GrabAcessTokenFromRequest(c.Request)
		refreshToken := utils.GrabAcessRefreshTokenFromRequest(c.Request)
		if token == "" || refreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token or refresh Token is empty"})
			return
		}
		crsfRefresh, _ := utils.GrabCrsfTokenFromRefreshToken(refreshToken)
		claims, _ := utils.GrabClaimsFromAcessToken(token)

		if claims.Crsf != crsfRefresh {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token and refresh token is not of the same user"})
		}
		crsf := utils.GenerateCrsfToken()
		user := domain.User{
			First_name: claims.First_name,
			Last_name:  claims.LastName,
			User_id:    claims.Uid,
			IsAdmin:    claims.IsAdmin,
		}
		acessToken, err := utils.GenerateToken(crsf, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		refreshToken, err = utils.GenerateRefreshToken(crsf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cookie := &http.Cookie{
			Name:     "token",
			Value:    acessToken,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		ok := controller.authService.UpdateRefreshToken(refreshToken, claims.Uid)
		if !ok {
			log.Print("Can not save the refresh token")
		}
		utils.SetCrsfTokenToClient(c.Writer, crsf)
		c.JSON(http.StatusOK, gin.H{"token": "acessToken"})
	}
}
