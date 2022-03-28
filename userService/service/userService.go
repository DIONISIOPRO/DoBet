package service

import (
	"github/namuethopro/dobet-user/domain"
	"github/namuethopro/dobet-user/repository"
)

type UserService interface {
	Users(page, perpage int64) ([]domain.User, error)
	GetUserById(userId string) (domain.User, error)
	GetUserByPhone(phone string) (domain.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewupUserService(userRepository repository.UserRepository) UserService{
	return &userService{
		repository:  userRepository,
	}
}



func (service *userService) Users(page, perpage int64) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Users(startIndex, perpage)
}

func (service *userService) GetUserById(userId string) (domain.User, error) {
	return service.repository.GetUserById(userId)
}

func (service *userService) GetUserByPhone(phone string) (domain.User, error) {
	return service.repository.GetUserByPhone(phone)
}
