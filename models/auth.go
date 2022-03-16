package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
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
	StandartClaims jwt.StandardClaims
}

type RefreshTokenClaims struct {
	CrsfToken      string
	StandartClaims jwt.StandardClaims
}

func (token TokenClaims) Valid() error {
	if token.StandartClaims.ExpiresAt >= time.Now().Unix() {
		return errors.New("token expires")
	}
	return nil
}

func (token RefreshTokenClaims) Valid() error {
	if token.StandartClaims.ExpiresAt >= time.Now().Unix() {
		return errors.New("token expires")
	}
	return nil
}
