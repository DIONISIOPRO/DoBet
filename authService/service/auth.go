package service

import (
	"encoding/json"
	"errors"

	"github/namuethopro/dobet-auth/domain"

	"github.com/streadway/amqp"
)

const (
	USERLOGIN  = "use.login"
	USERLOGOUT = "user.logout"
)

type (
	loginEvent struct {
		UserId string
	}
	logOutEvent struct {
		UserId string
	}
	AuthEventProcessor interface {
		AddUser(user domain.User) error
	}
	AuthEventSubscriber interface {
		SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
	}
	AuthEventPublisher interface {
		Publish(name string, body []byte) error
	}
	AuthEventListenner interface {
		ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte))
	}
	Authrepo interface {
		Login(phone string) (domain.User, error)
		SignUp(user domain.User) (string, error)
		AddRefreshToken(refreshtoken string) error
		GetRefreshTokens(userid string) ([]string, error)
	}
	PasswordVerifier interface {
		VerifyPassword(password, hash string) bool
	}

	authService struct {
		PasswordVerifier  PasswordVerifier
		authRepo          Authrepo
		LoginStateManager LogInStateManager
		jwtmanager        jwtManager
		eventManager      authEventManager
	}

	jwtManager interface {
		GenerateAcessToken(user domain.User) (string, error)
		GenerateRefreshToken(userid string) (string, error)
		VerifyToken(incomingtoken string) bool
		IsTokenExpired(token string) (bool, error)
		ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error)
	}
	authEventManager interface {
		AuthEventListenner
		AuthEventProcessor
		AuthEventPublisher
		AuthEventSubscriber
		CreateQueues([]string) error
	}
)

func NewAuthService(LoginStateManager *LogInStateManager, Authrepo Authrepo, eventManager authEventManager, jwtmanager jwtManager, PasswordVerifier PasswordVerifier) *authService {
	service := &authService{
		LoginStateManager: *LoginStateManager,
		PasswordVerifier:  PasswordVerifier,
		authRepo:          Authrepo,
		eventManager:      eventManager,
		jwtmanager:        jwtmanager}
	return service
}

func (service *authService) Login(user domain.LoginDetails) (token, refreshToken string, err error) {
	if user.Phone == "" || user.Password == "" {
		return "", "", errors.New("invalid user details")
	}
	localuser, err := service.authRepo.Login(user.Phone)
	if err != nil {
		return "", "", err
	}
	ok := service.PasswordVerifier.VerifyPassword(user.Password, localuser.Hashed_password)
	if !ok {
		return "", "", errors.New("user invalid")
	}
	token, err = service.jwtmanager.GenerateAcessToken(localuser)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.jwtmanager.GenerateRefreshToken(localuser.User_id)
	if err != nil {
		return "", "", err
	}
	if err != nil {
		return "", "", err
	}
	err = service.authRepo.AddRefreshToken(refreshToken)
	service.publishLoginEvent(localuser.User_id)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func (service *authService) Logout(token string) error {
	ok := service.jwtmanager.VerifyToken(token)
	if !ok {
		return errors.New("token invalid")
	}
	claims, err := service.jwtmanager.ExtractClaimsFromAcessToken(token)
	if err != nil {
		return err
	}
	userId := claims.Id
	err = service.publishLogOutEvent(userId)
	if err != nil {
		return err
	}
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
	if err != nil {
		return "", "", err
	}
	user := domain.User{
		First_name:   claims.First_name,
		Last_name:    claims.Last_name,
		Phone_number: claims.Phone,
		User_id:      claims.Id,
	}
	refreshtokens, err := service.authRepo.GetRefreshTokens(user.User_id)
	if err != nil {
		return "", "", err
	}
	if len(refreshtokens) <= 0 {
		return "", "", errors.New("invalid refresh token")
	}
	refreshCount := 0
	for _, rt := range refreshtokens {
		if rt == refreshToken {
			refreshCount++
		}
	}
	if refreshCount <= 0 {
		return "", "", errors.New("invalid refresh token")
	}
	acessToken, err = service.jwtmanager.GenerateAcessToken(user)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.jwtmanager.GenerateRefreshToken(claims.Id)
	if err != nil {
		return "", "", err
	}
	err = service.authRepo.AddRefreshToken(refreshToken)
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
