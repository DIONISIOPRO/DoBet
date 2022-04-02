package service

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github/namuethopro/dobet-user/domain"
	"github/namuethopro/dobet-user/event"
	"github/namuethopro/dobet-user/repository"

	"golang.org/x/crypto/bcrypt"
)

const (
	USERLOGIN           = "use.login"
	USERLOGOUT          = "user.logout"
	USERSIGNUP          = "user.signup"

)
type LoginEvent struct{
	UserId string
}
type LogOutEvent struct{
	UserId string
}
type SignUpEvent struct{
	User domain.User
}


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
	eventManager event.EventManager
}

func NewAuthService(authRepository repository.AuthRepository, logoutManager *LogoutStateManager, eventManager event.EventManager) AuthService {
	once := sync.Once{}
	service := &authService{
		authRepository: authRepository,
		LogoutManager:  *logoutManager,
		eventManager: eventManager,
	}
	once.Do(service.RegisterAdmin)
	err := service.eventManager.CreateQueues(QueuesToPublish)
	if err != nil{
		log.Fatal(err.Error())
	}
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
	userlogin := LoginEvent{
		UserId: userResponse.User_id,
	}
	data, err := json.Marshal(userlogin)
	if err != nil{
		return domain.User{}, err
	}
	err = service.eventManager.Publish(USERLOGIN, data)
	if err != nil{
		return domain.User{}, err
	}
	return userResponse, err
}

func (service *authService) Logout(userId string) {
	service.LogoutManager.Logout(userId)
	userlogout := LogOutEvent{
		UserId: userId,
	}
	data, err := json.Marshal(userlogout)
	if err != nil{
		log.Fatal(err)
		return
	}
	err = service.eventManager.Publish(USERLOGOUT, data)
	if err != nil{
		log.Fatal(err)
		return 
	}
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
	userEvent := SignUpEvent{
		User:user ,
	}
	data, err := json.Marshal(userEvent)
	if err != nil{
		return "", err
	}
	name, err := service.authRepository.SignUp(user)
	if err != nil{
		return name, err
	}
	err = service.eventManager.Publish(USERSIGNUP, data)
	if err != nil{
		return "", err
	}
	return name, nil
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
		First_name:   "Dionisio",
		Last_name:    "Namuetho",
		Phone_number: "878819968",
		Password:     "NAMUETHO",
		IsAdmin:      true,
	}
	user1 := domain.User{
		First_name:   "Dionisio",
		Last_name:    "Namuetho",
		Phone_number: "852798408",
		Password:     "NAMUETHO",
		IsAdmin:      true,
	}
	user.Hashed_password = service.HasPassword(user.Password)
	_, err := service.authRepository.SignUp(user)
	if err != nil {
	}
	user1.Hashed_password = service.HasPassword(user1.Password)

	_, err = service.authRepository.SignUp(user1)
}
