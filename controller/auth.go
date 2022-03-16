package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/service"
	"gitthub.com/dionisiopro/dobet/utils"

)

type authController struct {
	authService service.AuthService
}

func NewAuthController() *betController {
	return &betController{}
}

func (controller *authController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		userlogin := models.LoginDetails{}
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

		PasswordSameErr := CompareHashedPassword(user.Hashed_password, userlogin.Password)
		if PasswordSameErr != nil {
			msg := "Please provide an valid login credential"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		crsf := utils.GenerateCrsfToken()
		claims := models.TokenClaims{
			IsAdmin:   false,
			Phone:     userlogin.Phone,
			CrsfToken: crsf,
			StandartClaims: jwt.StandardClaims{
				Subject:   user.User_id,
				ExpiresAt: time.Now().Local().Add(time.Minute * 30).Unix(),
			},
		}

		acessToken, err := utils.GenerateNewAcessToken(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		refreshTokenClaims := models.RefreshTokenClaims{
			CrsfToken: crsf,
			StandartClaims: jwt.StandardClaims{
				Subject:   user.User_id,
				ExpiresAt: time.Now().Local().Add(time.Hour * 24 * 7).Unix(),
			},
		}

		refreshToken, err := utils.GenerateNewRefreshToken(refreshTokenClaims)
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
		c.JSON(http.StatusOK, gin.H{"token": acessToken})
		return
	}
}

func (controller *authController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		err := c.BindJSON(&user)
		if err != nil {
			msg := "Please povide a valid user"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
			return
		}
		user.Hashed_password, err = HasPassword(user.Password)
		if err != nil {
			msg := "Please povide a valid user"
			c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
			return
		}
		if err := controller.authService.SignUp(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, user.Phone_number)
		return

	}
}

func (controller *authController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutUser := models.LogoutDetails{}
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
		phone, err := utils.GrabPhoneFromAcessToken(acessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if phone != logoutUser.Phone {
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
		userId, _ := utils.GrabUuidFromAcessToken(acessToken)
		controller.authService.UpdateRefreshToken("", userId)
		c.JSON(http.StatusOK, logoutUser)
		return
	}
}

func (controller *authController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GrabAcessTokenFromRequest(c.Request)
		refreshToken := utils.GrabAcessRefreshTokenFromRequest(c.Request)
		if token == "" || refreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token or refresh Token is empty"})
			return
		}
		crsfAcess, _ := utils.GrabCrsfTokenFromAcessToken(token)
		crsfRefresh, _ := utils.GrabCrsfTokenFromRefreshToken(refreshToken)
		userid, _ := utils.GrabUuidFromAcessToken(token)
		phone, _:= utils.GrabPhoneFromAcessToken(token)

		if crsfAcess != crsfRefresh {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token and refresh token is not of the same user"})
		}
		
		crsf := utils.GenerateCrsfToken()
		claims := models.TokenClaims{
			IsAdmin:   false,
			Phone:     phone,
			CrsfToken: crsf,
			StandartClaims: jwt.StandardClaims{
				Subject:   userid,
				ExpiresAt: time.Now().Local().Add(time.Minute * 30).Unix(),
			},
		}

		acessToken, err := utils.GenerateNewAcessToken(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		refreshTokenClaims := models.RefreshTokenClaims{
			CrsfToken: crsf,
			StandartClaims: jwt.StandardClaims{
				Subject:   userid,
				ExpiresAt: time.Now().Local().Add(time.Hour * 24 * 7).Unix(),
			},
		}
		refreshToken, err = utils.GenerateNewRefreshToken(refreshTokenClaims)
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
		ok := controller.authService.UpdateRefreshToken(refreshToken, userid)
		if !ok {
			log.Print("Can not save the refresh token")
		}
		utils.SetCrsfTokenToClient(c.Writer, crsf)
		c.JSON(http.StatusOK, gin.H{"token": acessToken})
		return
	
	}
}
