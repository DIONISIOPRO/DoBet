package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	jwtManager interface {
		VerifyToken(incomingtoken string) bool
	}
)
type jwtMiddleWare struct {
	jwtmanager jwtManager
}

func NewjwtMiddleWare(jwtmanager jwtManager) *jwtMiddleWare {
	return &jwtMiddleWare{
		jwtmanager: jwtmanager,
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
		c.Next()
	}
}
