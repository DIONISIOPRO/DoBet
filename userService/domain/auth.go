package domain

import "github.com/golang-jwt/jwt"

type (
	LoginDetails struct {
		Phone    string `json:"phone_number"`
		Password string `json:"password"`
	}

	LogoutDetails struct {
		Phone    string `json:"phone_number"`
		Password string `json:"password"`
	}

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
