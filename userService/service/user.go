package service

import (
	"errors"
	"github/namuethopro/dobet-user/domain"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

const (
	USERDELETE          = "user.delete"
	USERCREATED         = "user.created"
	USERUPDATE          = "user.logout"
	USERCONFIRMWITHDRAW = "user.confirm.withdraw"
	USERCONFIRMBET      = "user.confirm.bet"
	USERREQUESTWITHDRAW = "user.request.withdraw"
	USERREQUESTBET      = "user.request.bet"
	USERDEPOSIT         = "user.deposit"
	USERWITHDRAW        = "user.withdraw"
	USERBET             = "user.bet"
	USERWIN             = "user.win"
)

var reserveMoney = make(map[string]struct {
	Amount float64
	Hash   string
})
var queuesToListenning = []string{
	USERREQUESTBET, USERREQUESTWITHDRAW, USERDEPOSIT, USERWIN, USERWITHDRAW, USERBET,
}
var queuesToPublish = []string{
	USERDELETE, USERUPDATE, USERCONFIRMWITHDRAW, USERCONFIRMBET, USERCREATED,
}

type (
	UserRepository interface {
		CreateUser(userRequest domain.User) (string, error)
		GetUserById(userId string) (domain.User, error)
		GetUserByPhone(phone string) (domain.User, error)
		GetUsers(startIndex, perpage int64) ([]domain.User, error)
		DeleteUser(userid string) error
		UpdateUser(userid string, user domain.User) error
		GetUserBalance(userId string) (float64, error)
		AddMoney(userId string, amount float64) error
		SubtractMoney(userId string, amount float64) error
	}
	userService struct {
		repository           UserRepository
		incomingEventHandler UserIncomingEventHandler
		eventManager         UserEventManager
		lock                 *sync.Mutex
	}
	UserEventManager interface {
		CreateQueues([]string) error
		SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
		Publish(name string, event domain.Event) error
		ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte) error)
	}
	UserIncomingEventHandler interface {
		SubtractBalance(data []byte) error
		CheckMoney(data []byte) error
		AddBalance(data []byte) error
	}
)

func NewUserService(userRepository UserRepository,
	eventManager UserEventManager,
	incomingEventHandler UserIncomingEventHandler,
	lock *sync.Mutex) *userService {
	userService := &userService{
		repository:           userRepository,
		eventManager:         eventManager,
		incomingEventHandler: incomingEventHandler,
		lock:                 lock,
	}
	return userService
}
func (service *userService) Start() {
	service.eventManager.CreateQueues(queuesToListenning)
	service.eventManager.CreateQueues(queuesToPublish)
	go service.listenningToqueue()
}

func (service *userService) CreateUser(user domain.User) (string, error) {
	err := user.Validate()
	if err != nil {
		return "", err
	}
	user.Account_balance = 0
	user.IsAdmin = false
	name, err := service.repository.CreateUser(user)
	if err != nil {
		return "", err
	}
	userCreated := &domain.UserCreatedEvent{
		User: user,
	}
	err = service.eventManager.Publish(USERCREATED, userCreated)
	return name, err
}

func (service *userService) GetUsers(page, perpage int64) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	users, err := service.repository.GetUsers(startIndex, perpage)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *userService) GetUserById(userId string) (domain.User, error) {
	if userId == "" {
		return domain.User{}, errors.New("user id is empty")
	}
	user, err := service.repository.GetUserById(userId)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (service *userService) GetUserByPhone(phone string) (domain.User, error) {
	if phone == "" {
		return domain.User{}, errors.New("user id is empty")
	}
	user, err := service.repository.GetUserByPhone(phone)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (service *userService) DeleteUser(userid string) error {
	if userid == "" {
		return errors.New("user id is empty")
	}
	err := service.repository.DeleteUser(userid)
	if err != nil {
		return err
	}
	userDeletedEvent := domain.UserDeletedEvent{
		UserId: userid,
	}
	err = service.eventManager.Publish(USERDELETE, userDeletedEvent)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) UpdateUser(userid string, user domain.User) error {
	if userid == "" {
		return errors.New("user id is empty")
	}
	err := service.repository.UpdateUser(userid, user)
	if err != nil {
		return err
	}
	userUpdateEvent := domain.UserUpdateEvent{
		User: user,
	}
	return service.eventManager.Publish(USERUPDATE, userUpdateEvent)
}

func (service *userService) listenningToqueue() {
	for _, queue := range queuesToListenning {
		topic, err := service.eventManager.SubscribeToQueue(queue)
		if err != nil {
			log.Print(err.Error())
		}
		switch queue {
		case USERDEPOSIT, USERWIN:
			service.eventManager.ListenningToqueue(topic, service.incomingEventHandler.AddBalance)
		case USERWITHDRAW, USERBET:
			service.eventManager.ListenningToqueue(topic, service.incomingEventHandler.SubtractBalance)
		case USERREQUESTBET, USERREQUESTWITHDRAW:
			service.eventManager.ListenningToqueue(topic, service.incomingEventHandler.CheckMoney)
		}

	}
}
