package controller

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET_KEY = os.Getenv("secretkey")

type SignedDetails struct {
	Phone      string
	Password   string
	First_name string
	LastName   string
	Uid        string
	jwt.StandardClaims
}

func HasPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func CompareHashedPassword(hashedPasword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}

func GenerateToken(phone, password, firstName, lastName, uid string) (string, error) {
	claims := &SignedDetails{
		First_name: firstName,
		LastName:   lastName,
		Phone:      phone,
		Password:   password,
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

func VerifyToken(token string) (SignedDetails, error) {
	localtoken, _ := jwt.ParseWithClaims(token, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	claims, ok := localtoken.Claims.(SignedDetails)
	if !ok {
		msg := "token is invalid"
		err := errors.New(msg)
		return SignedDetails{}, err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		msg := "token is expired"
		err := errors.New(msg)
		return SignedDetails{}, err
	}
	return claims, nil
}
