package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	User_id         string             `json:"user_id" bson:"user_id"`
	First_name      string             `json:"first_name" bson:"first_name" validate:"required"`
	Last_name       string             `json:"last_name" bson:"last_name" validate:"required"`
	Phone_number    string             `json:"phone_number" bson:"phone_number" validate:"required"`
	Account_balance float64            `json:"account_balance" bson:"account_balance"`
	Password        string             `json:"password" bson:"_"  validate:"required"`
	Created_at      time.Time          `json:"created_at" bson:"created_at"`
	Updated_at      time.Time          `json:"updated_at" bson:"updated_at"`
	IsAdmin         bool               `json:"is_admin" bson:"is_admin"`
	RefreshTokens   []string           `json:"refresh_tokens" bson:"refresh_tokens"`
	Hashed_password string             `json:"hashed_password" bson:"hashed_password"`
}

type UserResponse struct {
	User_id       string    `json:"user_id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Phone_number  string    `json:"phone_number"`
	Created_at    time.Time `json:"created_at"`
	IsAdmin       bool      `json:"is_admin"`
}

type UserLoginRequest struct {
	Phone_number string `json:"phone_number" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type UserSignUpRequest struct {
	First_name   string `json:"first_name" bson:"first_name" validate:"required"`
	Last_name    string `json:"last_name" bson:"last_name" validate:"required"`
	Phone_number string `json:"phone_number" bson:"phone_number" validate:"required"`
	Password     string `json:"password" bson:"_"  validate:"required"`
}

func (user User) ToResponse() UserResponse {
	responseUser := UserResponse{
		User_id: user.User_id,
		First_name: user.First_name,
		Last_name: user.Last_name,
		Phone_number: user.Phone_number,
		Created_at: user.Created_at,
		IsAdmin: user.IsAdmin,
	}
	return responseUser
}

func (user User) FromUserSignUp(signup UserSignUpRequest) *User{
	return &User{
		First_name: signup.First_name,
		Last_name: signup.Last_name,
		Phone_number: signup.Phone_number,
		Password: signup.Password,
	}

}