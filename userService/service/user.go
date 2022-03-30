package service

import (
	"errors"
	"github/namuethopro/dobet-user/domain"
	"github/namuethopro/dobet-user/repository"
)

type UserService interface {
	Users(page, perpage int64) ([]domain.UserResponse, error)
	GetUserById(userId string) (domain.UserResponse, error)
	GetUserByPhone(phone string) (domain.UserResponse, error)
	DeleteUser(userid string) error
	UpdateUser(userid string) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) Users(page, perpage int64) ([]domain.UserResponse, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	users, err := service.repository.Users(startIndex, perpage)

	if err != nil {
		return nil, err
	}
	userResponse := []domain.UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, user.ToResponse())
	}
	return userResponse, nil
}

func (service *userService) GetUserById(userId string) (domain.UserResponse, error) {
	user, err := service.repository.GetUserById(userId)
	if err != nil{
		return domain.UserResponse{}, err
	}
	return user.ToResponse(), nil
}

func (service *userService) GetUserByPhone(phone string) (domain.UserResponse, error) {
	user, err := service.repository.GetUserByPhone(phone)
	if err != nil{
		return domain.UserResponse{}, err
	}
	return user.ToResponse(), nil
}

func (service *userService) DeleteUser(userid string) error {
	if userid == "" {
		return errors.New("user id is empty")
	}
	return service.repository.DeleteUser(userid)
}

func (service *userService) UpdateUser(userid string) error {
	if userid == "" {
		return errors.New("user id is empty")
	}
	return service.repository.UpdateUser(userid)
}
