package middleware

import (
	"github/namuethopro/dobet-auth/domain"
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


