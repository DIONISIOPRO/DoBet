package domain

import "github.com/golang-jwt/jwt"

type (
	TokenClaims struct {
		Admin      bool
		First_name string
		Last_name  string
		Phone      string
		jwt.StandardClaims
	}

	RefreshTokenClaims struct {
		jwt.StandardClaims
	}
)