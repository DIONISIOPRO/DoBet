package domain

import (
	"errors"
	"strconv"
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

func (user *User) Validate() error {
	_, err := strconv.Atoi(user.Phone_number)
	if err != nil {
		return errors.New("your number is not valid")
	}
	if len([]rune(user.Phone_number)) != 9 {
		return errors.New("the lenght of number should be 9")
	}
	return nil
}

func (user *User) AddRefreshToken(refreshToken string) error {
	user.RefreshTokens = append(user.RefreshTokens, refreshToken)
	return user.Validate()
}

func (user *User) PromoteToAdmin() error {
	err := user.Validate()
	if err != nil{
		return err
	}
	user.IsAdmin = true
	return nil
}

func (user *User) Update() error {
	err := user.Validate()
	if err != nil{
		return err
	}
	user.Updated_at = time.Now().Local()
	return nil
}
