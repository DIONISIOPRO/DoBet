package service

import (
	"errors"
	"log"
	"time"

	"github/namuethopro/dobet-user/domain"
	 "github/namuethopro/dobet-user/repository"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Login(domain.LoginDetails) (domain.User, error)
	SignUp(user domain.User) error
	GetRefreshToken(userId string) ([]string, error)
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

func (service *authService) Login(user domain.LoginDetails) (domain.User, error) {
	if user.Phone == "" {
		return domain.User{}, errors.New("phone invalid")
	}
	return service.authRepository.Login(user.Phone)
}
func (service *authService) SignUp(user domain.User) error {
	validate := validator.New()
	err := validate.Struct(user)
	user.Account_balance = 0
	user.Created_at = time.Now().Local().Unix()
	if err != nil {
		log.Println(err)
		return err
	}
	return service.authRepository.SignUp(user)
}

func (service *authService) GetRefreshToken(userId string) ([]string, error) {
	if userId == "" {
		return nil, errors.New("id is invalid")
	}
	return service.authRepository.GetRefreshToken(userId)
}

func (service *authService) UpdateRefreshToken(refreshToken, userId string) bool {
	if refreshToken == "" || userId == ""{
		return false
	}
	return service.authRepository.UpdateRefreshToken(refreshToken, userId)
}
