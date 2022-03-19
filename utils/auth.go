package utils

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SECRETE = "SECRETE"

func GenerateCrsfToken() string {
	crsf := primitive.NewObjectID()
	return crsf.String()
}

func VerifyAcessToken(accesToken string) bool {
	token, _ := jwt.Parse(accesToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
	})
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		fmt.Print("assignint method error")
		return false
	}
	return true
}

func GenerateNewAcessToken(tokenClaims models.TokenClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString([]byte(SECRETE))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func GenerateNewRefreshToken(refresTokenClaims models.RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresTokenClaims)
	signed, err := token.SignedString([]byte(SECRETE))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func VerifyIsAdmin(acessToken string) bool {
	token, err := jwt.ParseWithClaims(acessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
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
		return []byte(SECRETE), nil
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
		return []byte(SECRETE), nil
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
		return []byte(SECRETE), nil
	})

	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.Subject, nil
}

func GrabPhoneFromAcessToken(acessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(acessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
	})

	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.Subject, nil
}

func GrabAcessTokenFromRequest(req *http.Request) string {
	fronHeader := req.Header.Get("token")
	return fronHeader
}

func GrabAcessRefreshTokenFromRequest(req *http.Request) string {
	return req.Header.Get("refreshtoken")
}
