package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	User_id         string             `json:"user_id" `
	First_name      string             `json:"first_name" validate:"required"`
	Last_name       string             `json:"last_name" validate:"required"`
	Phone_number    string             `json:"phone_number" validate:"required,min=9,max=9"`
	Account_balance float64            `json:"account_balance"`
	Password        string             `json:"password" validate:"required"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	RefreshToken    string             `json:"refresh_token" bson:"refresh_token"`
	Hashed_password string             `json:"hashed_password"`
}
