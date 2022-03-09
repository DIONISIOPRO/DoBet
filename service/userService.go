package service

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var UserService = &userService{}

type userService struct {
	repository repository.UserRepository
}

func (service *userService) SetupUserService(userRepository repository.UserRepository) {
	UserService.repository = userRepository
}

func (service *userService) Deposit(amount float64, userid string) error {
	return service.repository.Deposit(amount, userid)
}

func (service *userService) Withdraw(amount float64, userid string) error {
	return service.repository.Withdraw(amount, userid)
}

func (service *userService) Login(user models.User) (models.User, error) {
	return service.repository.Login(user)
}

func (service *userService) SignUp(user models.User) error {

	return service.repository.SignUp(user)

}

func (service *userService) Users(startIndex, perpage int64) ([]models.User, error) {
	return service.repository.Users(startIndex, perpage)
}
