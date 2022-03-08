package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var UserService userService

type userService struct {
	repo repositories.UserRepository
}

func (service *userService) SetupUserService(userRepository repositories.UserRepository) *userService {
	UserService.repo = userRepository
	return &UserService
}

func (service *userService) Deposit(amount float64, userid string) error {
	return service.repo.Deposit(amount, userid)
}

func (service *userService) Withdraw(amount float64, userid string) error {
	return service.repo.Withdraw(amount, userid)
}

func (service *userService) Login(user models.User) (models.User, error) {
	return service.repo.Login(user)
}

func (service *userService) SignUp(user models.User) error {

	return service.repo.SignUp(user)

}

func (service *userService) Users(startIndex, perpage int64) ([]models.User, error) {
	return service.repo.Users(startIndex, perpage)
}
