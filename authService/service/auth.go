package service

import (
	"errors"

	"github/namuethopro/dobet-auth/domain"

	"github.com/streadway/amqp"
)

type (
	AuthEventProcessor interface {
		AddUser([]byte) error
		RemoveUser([]byte) error
		UpdateUser(data []byte) error
	}
	AuthEventSubscriber interface {
		SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
	}
	AuthEventPublisher interface {
		Publish(name string, event domain.Event) error
	}

	AuthEventQueueCreator interface {
		CreateQueues([]string) error
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
		PasswordVerifier PasswordVerifier
		authRepo         Authrepo
		jwtmanager       jwtManager
		eventManager     authEventManager
	}

	jwtManager interface {
		GenerateAcessToken(user domain.User) (string, error)
		GenerateRefreshToken(userid string) (string, error)
		VerifyToken(incomingtoken string) bool
		IsTokenExpired(token string) (bool, error)
		ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error)
	}
	authEventManager interface {
		AuthEventProcessor
		AuthEventPublisher
		AuthEventSubscriber
		AuthEventQueueCreator
	}
)

func newAuthService(Authrepo Authrepo, eventManager authEventManager, jwtmanager jwtManager, PasswordVerifier PasswordVerifier) *authService {
	service := &authService{
		PasswordVerifier: PasswordVerifier,
		authRepo:         Authrepo,
		eventManager:     eventManager,
		jwtmanager:       jwtmanager}
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
	userlogin := domain.LoginEvent{
		UserId: id,
	}
	err := service.eventManager.Publish(domain.USERLOGIN, userlogin)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) publishLogOutEvent(id string) error {
	userlogout := domain.LogOutEvent{
		UserId: id,
	}

	err := service.eventManager.Publish(domain.USERLOGOUT, userlogout)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) StartEventHandler(done <-chan bool) {
	//creating queues wich i will publish
	err := service.eventManager.CreateQueues(domain.EventsToPublish)
	if err != nil {
		panic("cann`t create queues to publish events")
	}
	//subscribing in queues where i will listenning to
	for _, queue := range domain.EventsToListenning {
		channel, err := service.eventManager.SubscribeToQueue(queue)
		if err != nil {
			panic("cann`t subscribe to all channels")
		}
		switch queue {
		case domain.USERCREATED:
			go processMessage(channel, service.eventManager.AddUser, done)
		case domain.USERUPDATE:
			go processMessage(channel, service.eventManager.UpdateUser, done)
		case domain.USERDELETE:
			go processMessage(channel, service.eventManager.RemoveUser, done)
		default:
			continue
		}
	}
}

func processMessage(queue <-chan amqp.Delivery, processor func([]byte) error, done <-chan bool) {
	goroutinesCountChann := make(chan int, 5)
	for q := range queue {
		select {
		case <-done:
			return
		default:
			goroutinesCountChann <- 1
			go func(delivery amqp.Delivery) {
				data := delivery.Body
				err := processor(data)
				if err != nil {
					delivery.Ack(false)
				}
				<-goroutinesCountChann
			}(q)
		}

	}
}