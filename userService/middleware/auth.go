package middleware

import (
	"github/namuethopro/dobet-user/auth"
	"github/namuethopro/dobet-user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JWTMiddleWare interface {
	Authenticated() gin.HandlerFunc
	IsAdmin() gin.HandlerFunc
	IsOwner() gin.HandlerFunc
}
type JWTMiddleWareImp struct {
	jwtmanager    auth.JWTManager
	LogoutManager *service.LogoutStateManager
}

func NewJwtMiddleware(jwtmanager auth.JWTManager, logoutManager *service.LogoutStateManager) JWTMiddleWare {
	return &JWTMiddleWareImp{
		jwtmanager:    jwtmanager,
		LogoutManager: logoutManager,
	}
}
func (manager *JWTMiddleWareImp) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
		}
		isvalid := manager.jwtmanager.VerifyToken(token)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is ivalid"})
			c.Abort()
		}
		claims, _ := manager.jwtmanager.ExtractClaimsFromAcessToken(token)
		ok := manager.LogoutManager.IsLogIn(claims.Id)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "your logout, pleasse login"})
			c.Abort()
		}
		c.Next()
	}
}

func (manager *JWTMiddleWareImp) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
		}
		claims, err := manager.jwtmanager.ExtractClaimsFromAcessToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}
		if !claims.Admin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "only for admin"})
			c.Abort()
		}
		c.Next()
	}
}

func (manager *JWTMiddleWareImp) IsOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		IdParam := c.Param("id")
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
		}
		claims, err := manager.jwtmanager.ExtractClaimsFromAcessToken(token)
		if err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}
		if IdParam != claims.Id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
		}
		c.Next()
	}
}
