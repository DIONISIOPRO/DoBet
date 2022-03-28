package utils

import (
	"github/namuethopro/dobet-user/domain"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	Phone      string
	First_name string
	LastName   string
	Crsf       string
	Uid        string
	IsAdmin    bool
	jwt.StandardClaims
}

type RefreshTokenDetails struct {
	Crsf string
	jwt.StandardClaims
}

var JWT_SECRET_KEY = "SECRET"

func GenerateToken(csrsf string, user domain.User) (string, error) {
	claims := &TokenDetails{
		Crsf:       csrsf,
		First_name: user.First_name,
		LastName:   user.Last_name,
		Phone:      user.Phone_number,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateRefreshToken(crsf string) (string, error) {
	claims := &RefreshTokenDetails{
		Crsf: crsf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(token string) bool {
	localtoken, _ := jwt.ParseWithClaims(token, &TokenDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	claims, ok := localtoken.Claims.(TokenDetails)
	if !ok {
		return false
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return false
	}
	return true
}

func VerifyIfIsExpiredToken(token string) bool{
	localtoken, _ := jwt.ParseWithClaims(token, &TokenDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	claims, ok := localtoken.Claims.(TokenDetails)
	if !ok {
		return false
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return true
	}
	return false
}