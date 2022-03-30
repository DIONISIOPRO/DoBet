package auth

import (
	"errors"
	"github/namuethopro/dobet-user/domain"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager interface {
	GetTokenFromRequest(req *http.Request) string
	GenerateAcessToken(user domain.User) (string, error)
	GenerateRefreshToken(userid string) (string, error)
	VerifyToken(token string) bool
	IsTokenExpired(token string) bool
	ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error)
}

type JWTManagerImp struct {
	PrivateKey []byte
}

func NewJwtManager(PrivateKey []byte) JWTManager {
	return &JWTManagerImp{
		PrivateKey: PrivateKey,
	}
}


func (manager *JWTManagerImp) GenerateAcessToken(user domain.User) (string, error) {
	claims := &domain.TokenClaims{
		Admin:      user.IsAdmin,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Phone:      user.Phone_number,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(manager.PrivateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (manager JWTManagerImp) GenerateRefreshToken(userid string) (string, error) {
	claims := &domain.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        userid,
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(manager.PrivateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (manager *JWTManagerImp) VerifyToken(incomingtoken string) bool {
	token, _ := jwt.ParseWithClaims(incomingtoken, &domain.TokenClaims{}, func(t *jwt.Token) (interface{}, error){
		return manager.PrivateKey, nil
	})
	claims, ok := token.Claims.(*domain.TokenClaims)
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	return ok && !isTokenExpires && isHMACMethothod
}

func (manager *JWTManagerImp) IsTokenExpired(incomingtoken string) bool {
	token, _ := jwt.ParseWithClaims(incomingtoken, domain.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return manager.PrivateKey, nil
	})
	claims, ok := token.Claims.(*domain.TokenClaims)
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	return ok && isTokenExpires && isHMACMethothod
}

func (manager *JWTManagerImp) GetTokenFromRequest(req *http.Request) string {
	fronHeader := req.Header.Get("token")
	return fronHeader
}

func (manager *JWTManagerImp) ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error) {
	token, _ := jwt.ParseWithClaims(acessToken, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return manager.PrivateKey, nil
	})
	claims, ok := token.Claims.(*domain.TokenClaims)
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	if !ok && isTokenExpires && !isHMACMethothod {
		return domain.TokenClaims{}, errors.New("token is ivalid")
	}
	return *claims, nil
}
