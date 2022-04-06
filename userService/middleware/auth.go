package middleware

import (
	"github/namuethopro/dobet-user/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	loginStateManager interface {
		IsLogIn(id string) bool
	}
	
	jwtManager interface {
		GenerateAcessToken(user domain.User) (string, error)
		GenerateRefreshToken(userid string) (string, error)
		VerifyToken(incomingtoken string) bool
		IsTokenExpired(incomingtoken string) (bool, error)
		ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error)
	}
)
type jwtMiddleWare struct {
	jwtmanager   jwtManager
	logInManager loginStateManager
}

func NewjwtMiddleWare(jwtmanager jwtManager, logInManager loginStateManager) *jwtMiddleWare {
	return &jwtMiddleWare{
		jwtmanager:   jwtmanager,
		logInManager: logInManager,
	}
}
func (manager *jwtMiddleWare) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
			return
		}
		isvalid := manager.jwtmanager.VerifyToken(token)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}
		claims, _ := manager.jwtmanager.ExtractClaimsFromAcessToken(token)
		ok := manager.logInManager.IsLogIn(claims.Id)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "your logout, pleasse login"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (manager *jwtMiddleWare) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
			return
		}
		claims, err := manager.jwtmanager.ExtractClaimsFromAcessToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		if !claims.Admin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "only for admin"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (manager *jwtMiddleWare) IsOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		IdParam := c.Param("id")
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is empty",
			})
			c.Abort()
			return
		}
		claims, err := manager.jwtmanager.ExtractClaimsFromAcessToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Admin {
			c.Next()
			return
		}
		if IdParam != claims.Id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}
