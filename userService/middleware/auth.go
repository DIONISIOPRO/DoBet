package middleware

import (
	"github/namuethopro/dobet-user/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GrabAcessTokenFromRequest(c.Request)
		isvalid := utils.VerifyToken(token)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func OnlyForAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GrabAcessTokenFromRequest(c.Request)
		isAdmin := utils.VerifyIsAdmin(token)
		if !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
		}
		c.Next()

	}
}

func IfIdParamIsOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		IdParam := c.Param("id")
		token := utils.GrabAcessTokenFromRequest(c.Request)
		claims, _ := utils.GrabClaimsFromAcessToken(token)
		if IdParam != claims.Uid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
		}
		c.Next()
	}
}
