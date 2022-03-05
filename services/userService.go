package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

type userService struct{
	repo repositories.UserRepository
}
func NewUserService(repo repositories.UserRepository) UserService{
	return &userService{
		repo: repo,
	}
}
func (service *userService) Deposit(amount float64, userid string) error {
	return service.repo.Deposit(amount, userid)
}

func(service *userService)  Withdraw(amount float64, userid string) error {
	return service.repo.Withdraw(amount, userid)
}


func(service *userService)  Login(user models.User) error {
	return service.repo.Login(user)
}

func(service *userService)  SignUp(user models.User) error {

	return service.repo.SignUp(user)

}

func (service *userService) Users() ([]models.User, error){
	return service.repo.Users()
}
