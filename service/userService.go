package service

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

type UserService interface {
	Users(page, perpage int64) ([]models.User, error)
	GetUserById(userId string) (models.User, error)
	GetUserByPhone(phone string) (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func SetupUserService(userRepository repository.UserRepository) UserService{
	return &userService{
		repository:  userRepository,
	}
}



func (service *userService) Users(page, perpage int64) ([]models.User, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Users(startIndex, perpage)
}

func (service *userService) GetUserById(userId string) (models.User, error) {
	return service.repository.GetUserById(userId)
}

func (service *userService) GetUserByPhone(phone string) (models.User, error) {
	return service.repository.GetUserByPhone(phone)
}
