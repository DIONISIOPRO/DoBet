package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateCrsfToken() string {
	crsf := primitive.NewObjectID()
	return crsf.String()
}

func VerifyAcessToken(accesToken string) bool {
	token, err := jwt.ParseWithClaims(accesToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})

	if err != nil {
		return false
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return false
	}
	_, ok = token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return false
	}
	if tokenclaims.StandartClaims.ExpiresAt < time.Now().Local().Unix() {
		return false
	}
	return true
}

func GenerateNewAcessToken(tokenClaims models.TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), tokenClaims)
	signed, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func GenerateNewRefreshToken(refresTokenClaims models.RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), refresTokenClaims)
	signed, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func VerifyIsAdmin(acessToken string) bool {
	token, err := jwt.ParseWithClaims(acessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		return false
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return false
	}
	return tokenclaims.IsAdmin
}

func GrabCrsfTokenFromAcessToken(acesstoken string) (string, error) {
	token, err := jwt.ParseWithClaims(acesstoken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.CrsfToken, nil
}

func GrabCrsfTokenFromRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*models.RefreshTokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.CrsfToken, nil
}

func GrabCrsfTokenFromRequest(req *http.Request) string {
	crsf := req.FormValue("X-CRSF-TOKEN")
	if crsf != "" {
		return crsf
	}
	return req.Header.Get("X-CRSF-TOKEN")
}

func SetCrsfTokenToClient(w http.ResponseWriter, crsf string) {
	w.Header().Set("X-CRSF-TOKEN", crsf)
}

func GrabUuidFromAcessToken(acessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(acessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})

	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.StandartClaims.Subject, nil
}

func GrabPhoneFromAcessToken(acessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(acessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})

	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.StandartClaims.Subject, nil
}

func GrabAcessTokenFromRequest(req *http.Request) string {
	cookie, err := req.Cookie("token")
	if err != nil || cookie.Value == "" {
		token := strings.Split(req.Header["Token"][0], "")
		if len(token) == 2 {
			return token[1]
		} else if len(token) == 1 {
			return token[0]
		} else {
			return ""
		}
	}
	return cookie.Value
}

func GrabAcessRefreshTokenFromRequest(req *http.Request) string {
	cookie, err := req.Cookie("refreshtoken")
	if err != nil || cookie.Value == "" {
		token := strings.Split(req.Header["Refreshtoken"][0], "")
		if len(token) == 2 {
			return token[1]
		} else if len(token) == 1 {
			return token[0]
		} else {
			return ""
		}
	}
	return cookie.Value
}
