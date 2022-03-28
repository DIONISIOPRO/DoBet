package utils

import (
	"github/namuethopro/dobet-user/domain"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SECRETE = "SECRETE"

func GenerateCrsfToken() string {
	crsf := primitive.NewObjectID().Hex()
	return crsf
}

func VerifyIsAdmin(acessToken string) bool {
	token, err := jwt.ParseWithClaims(acessToken, &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
	})
	if err != nil {
		return false
	}
	tokenclaims, ok := token.Claims.(TokenDetails)
	if !ok {
		return false
	}
	return tokenclaims.IsAdmin
}

func GrabCrsfTokenFromAcessToken(acesstoken string) (string, error) {
	token, err := jwt.ParseWithClaims(acesstoken, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
	})
	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*domain.TokenClaims)
	if !ok {
		return "", err
	}
	return tokenclaims.CrsfToken, nil
}

func GrabCrsfTokenFromRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
	})
	if err != nil {
		return "", err
	}
	tokenclaims, ok := token.Claims.(*domain.RefreshTokenClaims)
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

func GrabClaimsFromAcessToken(acessToken string) (TokenDetails, error) {
	token, err := jwt.ParseWithClaims(acessToken, &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETE), nil
	})

	if err != nil {
		return TokenDetails{}, err
	}
	tokenclaims, ok := token.Claims.(TokenDetails)
	if !ok {
		return TokenDetails{}, err
	}
	return tokenclaims, nil
}

func GrabAcessTokenFromRequest(req *http.Request) string {
	fronHeader := req.Header.Get("token")
	return fronHeader
}

func GrabAcessRefreshTokenFromRequest(req *http.Request) string {
	return req.Header.Get("refreshtoken")
}
