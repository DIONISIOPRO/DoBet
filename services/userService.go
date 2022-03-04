package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func Deposit(amount float64, userid string) error {
	return repositories.Deposit(amount, userid)
}

func Withdraw(amount float64, userid string) error {
	return nil
}


func Login(user models.User) error {
	return repositories.Login(user)
}

func SignUp(user models.User) error {

	return repositories.SignUp(user)

}
