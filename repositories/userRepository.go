package repositories

import (
	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
)

var userColletion = database.OpenCollection("users")

func Deposit(amount float64, userid string) error {

	return nil

}

func Withdraw(amount float64, userid string) error {
	return nil
}


func Login(user models.User) error {
	return nil
}

func SignUp(user models.User) error {

	return nil

}

func Win(amount float64, user_id string) {
	
}

