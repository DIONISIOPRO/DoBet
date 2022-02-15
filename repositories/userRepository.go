package repositories

import "gitthub.com/dionisiopro/dobet/models"

func Deposit(amount float64, user models.User) error {

	return nil

}

func Withdraw(amount float64, userAccount models.User) error {
	return nil
}


func Login(user models.User) error {
	return nil
}

func SignUp(user models.User) error {

	return nil

}

func Win(amount float64, user models.User) {
	user.Account_balance += amount
}
