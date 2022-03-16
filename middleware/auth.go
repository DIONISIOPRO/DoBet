package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/utils"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GrabAcessTokenFromRequest(c.Request)
		isvalid := utils.VerifyAcessToken(token)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
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
		IdRequest, _ := utils.GrabUuidFromAcessToken(token)
		if IdParam != IdRequest {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not acess this resource"})
			c.Abort()
		}
		c.Next()
	}
}
