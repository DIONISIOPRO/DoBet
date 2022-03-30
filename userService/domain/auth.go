package domain

import "github.com/golang-jwt/jwt"

type LoginDetails struct {
	Phone    string `json:"phone_number"`
	Password string `json:"password"`
}

type LogoutDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type TokenClaims struct {
	Admin      bool
	First_name string
	Last_name  string
	Phone      string
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
}

