package service

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github/namuethopro/dobet-user/domain"

	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
)

const (
	USERLOGIN  = "use.login"
	USERLOGOUT = "user.logout"
	USERSIGNUP = "user.signup"
)

type (
	loginEvent struct {
		UserId string
	}
	logOutEvent struct {
		UserId string
	}
	signUpEvent struct {
		User domain.User
	}

	authService struct {
		authRepository    authRepository
		LoginStateManager LogInStateManager
		jwtmanager        jwtManager
		eventManager      authEventManager
	}

	authRepository interface {
		Login(phone string) (domain.User, error)
		SignUp(user domain.User) (string, error)
		GetRefreshTokens(userId string) ([]string, error)
		AddRefreshToken(refreshToken, userId string) error
		RevokeRefreshToken(userId string) error
		CleanUpDataBase() error
	}
	jwtManager interface {
		GenerateAcessToken(user domain.User) (string, error)
		GenerateRefreshToken(userid string) (string, error)
		IsTokenExpired(token string) (bool, error)
		ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error)
	}
	authEventManager interface {
		CreateQueues([]string) error
		SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
		Publish(name string, body []byte) error
		ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte))
	}
)

func NewAuthService(authRepository authRepository,
	LoginStateManager *LogInStateManager, eventManager authEventManager,
	jwtmanager jwtManager) *authService {
	once := sync.Once{}
	service := &authService{
		authRepository:    authRepository,
		LoginStateManager: *LoginStateManager,
		eventManager:      eventManager,
		jwtmanager:        jwtmanager,
	}
	once.Do(service.RegisterAdmin)
	err := service.eventManager.CreateQueues(queuesToPublish)
	if err != nil {
		log.Fatal(err.Error())
	}
	return service
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
	user.Hashed_password = hasFromPassword(user.Password)
	name, err := service.authRepository.SignUp(user)
	if err != nil {
		return "", err
	}
	service.publishSignUpEvent(user)
	return name, nil
}

func (service *authService) Login(user domain.LoginDetails) (token, refreshToken string, err error) {
	if user.Phone == "" || user.Password == "" {
		return "", "", errors.New("invalid user details")
	}
	userResponse, err := service.authRepository.Login(user.Phone)
	if err != nil {
		return "", "", err
	}
	ok := verifyPassword(userResponse.Hashed_password, user.Password)
	if !ok {
		return "", "", errors.New("user invalid")
	}
	token, err = service.jwtmanager.GenerateAcessToken(userResponse)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.jwtmanager.GenerateRefreshToken(userResponse.User_id)
	if err != nil {
		return "", "", err
	}
	err = service.addRefreTokenToUser(refreshToken, userResponse.User_id)
	if err != nil {
		return "", "", err
	}
	service.LoginStateManager.LogIn(userResponse.User_id)
	err = service.publishLoginEvent(userResponse.User_id)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func (service *authService) Logout(token string) error {
	claims, err := service.jwtmanager.ExtractClaimsFromAcessToken(token)
	if err != nil {
		return err
	}
	userId := claims.Id
	err = service.publishLogOutEvent(userId)
	if err != nil {
		return err
	}
	err = service.revokeRefreshTokens(userId)
	if err != nil {
		return err
	}
	service.LoginStateManager.Logout(userId)
	return nil
}

func (service *authService) RefreshToken(token string) (acessToken, refreshToken string, err error) {
	ok, err := service.jwtmanager.IsTokenExpired(token)
	if err != nil || !ok {
		return "", "", errors.New("can not refresh with this your token")
	}
	claims, err := service.jwtmanager.ExtractClaimsFromAcessToken(token)
	if err != nil {
		return "", "", err
	}
	tokens, err := service.getRefreshTokens(claims.Id)
	if err != nil {
		return "", "", err
	}
	if len(tokens) == 0 {
		return "", "", errors.New("your need login")
	}
	user := domain.User{
		First_name: claims.First_name,
		Last_name:  claims.Last_name,
		User_id:    claims.Id,
		IsAdmin:    claims.Admin,
	}
	acessToken, err = service.jwtmanager.GenerateAcessToken(user)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.jwtmanager.GenerateRefreshToken(claims.Id)
	if err != nil {
		return "", "", err
	}
	err = service.addRefreTokenToUser(refreshToken, claims.Id)
	if err != nil {
		return "", "", err
	}
	return acessToken, refreshToken, nil
}

func (service *authService) publishLoginEvent(id string) error {
	userlogin := loginEvent{
		UserId: id,
	}
	data, err := json.Marshal(userlogin)
	if err != nil {
		return err
	}
	err = service.eventManager.Publish(USERLOGIN, data)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) publishLogOutEvent(id string) error {
	userlogout := logOutEvent{
		UserId: id,
	}
	data, err := json.Marshal(userlogout)
	if err != nil {
		return err
	}
	err = service.eventManager.Publish(USERLOGOUT, data)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) publishSignUpEvent(user domain.User) error {
	userEvent := signUpEvent{
		User: user,
	}
	data, err := json.Marshal(userEvent)
	if err != nil {
		return err
	}
	err = service.eventManager.Publish(USERSIGNUP, data)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) getRefreshTokens(userId string) ([]string, error) {
	if userId == "" {
		return nil, errors.New("id is invalid")
	}
	return service.authRepository.GetRefreshTokens(userId)
}

func (service *authService) revokeRefreshTokens(userId string) error {
	if userId == "" {
		return errors.New("user id is ivalid")
	}
	return service.authRepository.RevokeRefreshToken(userId)
}

func (service *authService) addRefreTokenToUser(refreshtoken, userId string) error {
	if userId == "" {
		return errors.New("user id is ivalid")
	}
	return service.authRepository.AddRefreshToken(refreshtoken, userId)
}

func hasFromPassword(password string) string {
	data, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(data)
}

func verifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *authService) RegisterAdmin() {
	service.authRepository.CleanUpDataBase()
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
	user.Hashed_password = hasFromPassword(user.Password)
	_, err := service.authRepository.SignUp(user)
	if err != nil {
		log.Print("error creating adm")
	}
	user1.Hashed_password = hasFromPassword(user1.Password)

	_, err = service.authRepository.SignUp(user1)
	if err != nil {
		log.Print("error creating adm")
	}
}
