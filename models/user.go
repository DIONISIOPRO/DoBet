package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	User_id         string             `json:"user_id" `
	First_name      string             `json:"first_name" validate:"required"`
	Last_name       string             `json:"last_name" validate:"required"`
	Phone_number    string             `json:"phone_number" validate:"required"`
	Account_balance float64            `json:"account_balance"`
	Password        string             `json:"password" validate:"required"`
	Created_at      int64              `json:"created_at"`
	Updated_at      int64              `json:"updated_at"`
	IsAdmin         bool               `json:"is_admin" validate:"required"`
	RefreshToken    string             `json:"refresh_token" bson:"refresh_token"`
	Hashed_password string             `json:"hashed_password"`
}
