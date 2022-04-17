package domain

import (
	"github.com/golang-jwt/jwt"
)

type (
	User struct {
		User_id         string   `json:"user_id" bson:"user_id"`
		First_name      string   `json:"first_name" bson:"first_name" validate:"required"`
		Last_name       string   `json:"last_name" bson:"last_name" validate:"required"`
		Phone_number    string   `json:"phone_number" bson:"phone_number" validate:"required"`
		Password        string   `json:"password" bson:"_"  validate:"required"`
		RefreshTokens   []string `json:"refresh_tokens" bson:"refresh_tokens"`
		Hashed_password string   `json:"hashed_password" bson:"hashed_password"`
	}
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
