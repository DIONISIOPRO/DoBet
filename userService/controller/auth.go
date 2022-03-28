package controller

import (
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

		Err := utils.CompareHashedPassword(user.Hashed_password, userlogin.Password)
		if Err != nil {
			msg := "Please provide an valid login credential"
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		token, refresh, crsf := GenerateNewTokens(user)
		SetTokenCookie(c, token)
		utils.SetCrsfTokenToClient(c.Writer, crsf)
		err = controller.authService.AddRefreToken(refresh, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
		RevokeTokenCookie(c)
		controller.authService.RevokeRefreToken(claims.Uid)
		c.JSON(http.StatusOK, logoutUser)
	}
}

func (controller *AuthController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GrabAcessTokenFromRequest(c.Request)
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token or refresh Token is empty"})
			return
		}
		ok := utils.VerifyIfIsExpiredToken(token)
		if !ok {
			c.Abort()
			return
		}
		claims, _ := utils.GrabClaimsFromAcessToken(token)
		tokens, err :=controller.authService.GetRefreshTokens(claims.Uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tokenExists := false
		for _, v := range tokens {
			if token == v{
				tokenExists = true
				break
			}
		}
		if !tokenExists{
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Token invalid"})
			return
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
		refreshToken, err := utils.GenerateRefreshToken(crsf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		SetTokenCookie(c, acessToken)
		Err := controller.authService.AddRefreToken(refreshToken, claims.Uid)
		if Err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		utils.SetCrsfTokenToClient(c.Writer, crsf)
		c.JSON(http.StatusOK, gin.H{"token": "acessToken"})
	}
}

func SetTokenCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}

func RevokeTokenCookie(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Local().Add(-100 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}

func GenerateNewTokens(user domain.User) (acessToken, refreshToken, crsf string) {
	crsf = utils.GenerateCrsfToken()
	acessToken, err := utils.GenerateToken(crsf, user)
	if err != nil {
		return "", "", ""
	}
	refreshToken, err = utils.GenerateRefreshToken(crsf)
	if err != nil {
		return "", "", ""
	}
	return
}
