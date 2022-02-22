package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func Deposit(amount float64, user models.User) error {
	return repositories.Deposit(amount, user)
}

func Withdraw(amount float64, userAccount models.User) error {
	return nil
}


func Login(user models.User) error {
	return repositories.Login(user)
}

func SignUp(user models.User) error {

	return repositories.SignUp(user)

}

func Win(amount float64, user models.User) {

	repositories.Win(amount, user)
}
