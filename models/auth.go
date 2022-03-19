package models

import (

	"github.com/dgrijalva/jwt-go"
)

type LoginDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LogoutDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type TokenClaims struct {
	IsAdmin        bool
	Phone          string
	CrsfToken      string
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	CrsfToken      string
	jwt.StandardClaims
}

// func (token TokenClaims) Valid() error {
// 	if token.StandartClaims.ExpiresAt >= time.Now().Unix() {
// 		return errors.New("token expires")
// 	}
// 	return nil
// }

// func (token RefreshTokenClaims) Valid() error {
// 	if token.StandartClaims.ExpiresAt >= time.Now().Unix() {
// 		return errors.New("token expires")
// 	}
// 	return nil
// }
