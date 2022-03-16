package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

type AuthService interface {
	Login(models.LoginDetails) (models.User, error)
	SignUp(user models.User) error
	GetRefreshToken(userId string) (string, error)
	UpdateRefreshToken(refreshToken, userId string) bool

}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (service *authService) Login(user models.LoginDetails) (models.User, error) {
	if user.Phone == "" {
		return models.User{}, errors.New("phone invalid")
	}
	return service.authRepository.Login(user.Phone)
}
func (service *authService) SignUp(user models.User) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return service.authRepository.SignUp(user)
}

func (service *authService) GetRefreshToken(userId string) (string, error) {
	if userId != "" {
		return service.authRepository.GetRefreshToken(userId)
	}
	return "", errors.New("id is invalid")
}

func (service *authService) UpdateRefreshToken(refreshToken, userId string) bool {
	return service.authRepository.UpdateRefreshToken(refreshToken, userId)
}
