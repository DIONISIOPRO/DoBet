package middleware

import (
	"errors"
	"github/namuethopro/dobet-user/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type (
	LoginStateManager interface {
		IsLogIn(id string) bool
	}
)
type jwtMiddleWare struct {
	PrivateKey   []byte
	logInManager LoginStateManager
}

func NewjwtMiddleWare(logInManager LoginStateManager, privateKey []byte) *jwtMiddleWare {
	return &jwtMiddleWare{
		logInManager: logInManager,
		PrivateKey:   privateKey,
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
		isvalid := verifyToken(token, manager.PrivateKey)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}
		claims, _ := extractClaimsFromAcessToken(token, manager.PrivateKey)
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
		claims, err := extractClaimsFromAcessToken(token, manager.PrivateKey)
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
		claims, err := extractClaimsFromAcessToken(token, manager.PrivateKey)
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

func verifyToken(incomingtoken string, privateKey []byte) bool {
	token, err := jwt.ParseWithClaims(incomingtoken, &domain.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err != nil {
		return false
	}
	claims, ok := token.Claims.(*domain.TokenClaims)
	if !ok {
		if err != nil {
			return false
		}
	}
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	if isTokenExpires {
		return false
	}
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	return isHMACMethothod
}

func extractClaimsFromAcessToken(acessToken string, privateKey []byte) (domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(acessToken, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err != nil {
		return domain.TokenClaims{}, errors.New("token invalid")
	}
	claims, ok := token.Claims.(*domain.TokenClaims)
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	if !ok && isTokenExpires && !isHMACMethothod {
		return domain.TokenClaims{}, errors.New("token is invalid")
	}
	return *claims, nil
}
