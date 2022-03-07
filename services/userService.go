package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var userRepository repositories.UserRepository

func NewUserService(userRepository repositories.UserRepository){
	userRepository = userRepository
}


func  Deposit(amount float64, userid string) error {
	return userRepository.Deposit(amount, userid)
}

func  Withdraw(amount float64, userid string) error {
	return userRepository.Withdraw(amount, userid)
}


func  Login(user models.User) (models.User,error) {
	return userRepository.Login(user)
}

func  SignUp(user models.User) error {

	return userRepository.SignUp(user)

}

func  Users(startIndex, perpage int64) ([]models.User, error){
	return userRepository.Users(startIndex, perpage)
}
