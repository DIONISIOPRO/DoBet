package service

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github/namuethopro/dobet-user/domain"

	"github/namuethopro/dobet-user/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(domain.LoginDetails) (domain.User, error)
	Logout(userId string)
	SignUp(userRequest domain.UserSignUpRequest) (string, error)
	GetRefreshTokens(userId string) ([]string, error)
	RevokeRefreToken(userId string) error
	AddRefreToken(refreshtoken, userId string) error
	HasPassword(password string) string
	VerifyPassword(hash, password string) bool
}

type authService struct {
	authRepository repository.AuthRepository
	LogoutManager  LogoutStateManager
}

func NewAuthService(authRepository repository.AuthRepository, logoutManager *LogoutStateManager) AuthService {
	once := sync.Once{}
	service := &authService{
		authRepository: authRepository,
		LogoutManager:  *logoutManager,
	}
	once.Do(service.RegisterAdmin)
	return service
}

func (service *authService) Login(user domain.LoginDetails) (domain.User, error) {
	if user.Phone == "" {
		return domain.User{}, errors.New("phone invalid")
	}
	userResponse, err := service.authRepository.Login(user.Phone)
	if err != nil {
		return domain.User{}, err
	}
	service.LogoutManager.LogIn(userResponse.User_id)
	return userResponse, err
}

func (service *authService) Logout(userId string) {
	service.LogoutManager.Logout(userId)
	service.RevokeRefreToken(userId)
}

func (service *authService) SignUp(userRequest domain.UserSignUpRequest) (string, error) {
	user := domain.User{}
	user = *user.FromUserSignUp(userRequest)
	if len(user.Phone_number) != 9 {
		return "", errors.New("the lenght of phone number should be 9")
	}
	user.Account_balance = 0
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.IsAdmin = false
	user.Hashed_password = service.HasPassword(user.Password)
	return service.authRepository.SignUp(user)
}

func (service *authService) GetRefreshTokens(userId string) ([]string, error) {
	if userId == "" {
		return nil, errors.New("id is invalid")
	}
	return service.authRepository.GetRefreshTokens(userId)
}

func (service *authService) RevokeRefreToken(userId string) error {
	if userId == "" {
		return errors.New("user id is ivalid")
	}
	return service.authRepository.RevokeRefreshToken(userId)
}

func (service *authService) AddRefreToken(refreshtoken, userId string) error {
	if userId == "" {
		return errors.New("user id is ivalid")
	}
	return service.authRepository.AddRefreshToken(refreshtoken, userId)
}

func (servce *authService) HasPassword(password string) string {
	data, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(data)
}

func (servce *authService) VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *authService) RegisterAdmin() {
	service.authRepository.CleanDataBase()
	user := domain.User{
		First_name:   "Dionisio Paulo",
		Last_name:    "Namuetho",
		Phone_number: "878819968",
		Password:     "NAMUETHO",
		IsAdmin:      true,
	}
	user1 := domain.User{
		First_name:   "Dionisio Paulo",
		Last_name:    "Namuetho",
		Phone_number: "852798408",
		Password:     "NAMUETHO",
		IsAdmin:      true,
	}
	user.Hashed_password = service.HasPassword(user.Password)
	userid, err := service.authRepository.SignUp(user)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(userid)
	user1.Hashed_password = service.HasPassword(user1.Password)

	userId1, err := service.authRepository.SignUp(user1)
	log.Print(userId1)

	if err != nil {
		fmt.Print(err)
	}
}
