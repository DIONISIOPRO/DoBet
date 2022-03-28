package domain

import "github.com/dgrijalva/jwt-go"

type LoginDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LogoutDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type TokenClaims struct {
	Admin        bool
	Phone          string
	CrsfToken      string
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	CrsfToken      string
	jwt.StandardClaims
}

