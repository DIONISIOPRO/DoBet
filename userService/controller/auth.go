package controller

import (
	"fmt"
	"github/namuethopro/dobet-user/auth"
	"github/namuethopro/dobet-user/domain"
	"github/namuethopro/dobet-user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	InvalidCredencialsErr = "invalid credential"
	TokenInvalidErr       = "token invalid"
)

type AuthController struct {
	authService service.AuthService
	jwtmanager  auth.JWTManager
}

func NewAuthController(
	authService service.AuthService, jwtmanager auth.JWTManager) *AuthController {
	return &AuthController{
		authService: authService,
		jwtmanager:  jwtmanager,
	}
}

func (controller *AuthController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtmanager := controller.jwtmanager
		authService := controller.authService
		userlogin := domain.LoginDetails{}
		err := c.BindJSON(&userlogin)
		checkBadRequestErr(c, err, InvalidCredencialsErr)
		user, err := authService.Login(userlogin)
		checkInternalServerErr(c, err)
		OK := controller.authService.VerifyPassword(user.Hashed_password, userlogin.Password)
		if !OK {
			c.JSON(http.StatusBadRequest, gin.H{"error": InvalidCredencialsErr})
			return
		}
		token, err := jwtmanager.GenerateAcessToken(user)
		checkInternalServerErr(c, err)
		refreshToken, err := jwtmanager.GenerateRefreshToken(user.User_id)
		checkInternalServerErr(c, err)
		err = controller.authService.AddRefreToken(refreshToken, user.User_id)
		checkInternalServerErr(c, err)
		c.JSON(http.StatusAccepted, gin.H{
			"token":        token,
			"refreshToken": refreshToken,
		})

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

func (controller *AuthController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		logoutUser := domain.LogoutDetails{}
		err := c.BindJSON(&logoutUser)
		checkBadRequestErr(c, err, InvalidCredencialsErr)
		acessToken := c.Request.Header.Get("token")
		if acessToken == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token iss empty"})
			return
		}
		claims, err := controller.jwtmanager.ExtractClaimsFromAcessToken(acessToken)
		checkUnauthorizedErr(c, err)
		if claims.Phone != logoutUser.Phone {
			c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidCredencialsErr})
			return
		}
		controller.authService.Logout(claims.Id)
		c.JSON(http.StatusOK, logoutUser)
	}
}

func (controller *AuthController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtmanager := controller.jwtmanager
		authService := controller.authService
		acessToken := c.Request.Header.Get("token")
		if acessToken == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
			return
		}
		ok, err := jwtmanager.IsTokenExpired(acessToken)
		checkInternalServerErr(c, err)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "your token is valid"})
			return
		}
		fmt.Print("esteve aqui aqui")
		claims, _ := jwtmanager.ExtractClaimsFromAcessToken(acessToken)
		tokens, err := authService.GetRefreshTokens(claims.Id)
		checkInternalServerErr(c, err)
		if len(tokens) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": TokenInvalidErr})
			return
		}
		user := domain.User{
			First_name: claims.First_name,
			Last_name:  claims.Last_name,
			User_id:    claims.Id,
			IsAdmin:    claims.Admin,
		}
		fmt.Print("chega aqui")
		acessToken, err = jwtmanager.GenerateAcessToken(user)
		checkInternalServerErr(c, err)
		refreshToken, err := jwtmanager.GenerateRefreshToken(claims.Id)
		checkInternalServerErr(c, err)
		err = authService.AddRefreToken(refreshToken, claims.Id)
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

func checkUnauthorizedErr(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
}

func checkBadRequestErr(c *gin.Context, err error, msg string) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": msg})
		return
	}
}
