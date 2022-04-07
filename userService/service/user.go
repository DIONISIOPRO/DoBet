package service

import (
	"encoding/json"
	"errors"
	"github/namuethopro/dobet-user/domain"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
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
	USERDELETE, USERUPDATE, USERCONFIRMWITHDRAW, USERCONFIRMBET,
}

type (
	userRepository interface {
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
		repository   userRepository
		eventManager userEventManager
		lock         *sync.Mutex
	}
	userEventManager interface {
		CreateQueues([]string) error
		SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
		Publish(name string, event interface{}) error
		ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte) error)
	}
)

func NewUserService(userRepository userRepository,
	eventManager userEventManager,
	lock *sync.Mutex) *userService {
	userService := &userService{
		repository:   userRepository,
		eventManager: eventManager,
		lock:         lock,
	}
	go func() {
		userService.eventManager.CreateQueues(queuesToPublish)
		userService.eventManager.CreateQueues(queuesToListenning)
	}()
	return userService
}
func (service *userService) CreateUser(userRequest domain.User) (string, error) {
	user := domain.User{}
	err := user.Validate()
	if err != nil {
		return "", err
	}
	user.Account_balance = 0
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.IsAdmin = false
	user.Hashed_password = hasFromPassword(user.Password)
	name, err := service.repository.CreateUser(user)
	if err != nil {
		return "", err
	}
	userCreated := domain.UserCreatedEvent{}
	service.eventManager.Publish(USERCREATED, userCreated)
	return name, nil
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
	user, err := service.repository.GetUserById(userId)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (service *userService) GetUserByPhone(phone string) (domain.User, error) {
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
	userDeletedEvent := &domain.UserDeletedEvent{
		UserId: userid,
	}
	return service.eventManager.Publish(USERDELETE, userDeletedEvent)
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

func (service *userService) ListenningToqueue() {
	for _, queue := range queuesToListenning {
		topic, err := service.eventManager.SubscribeToQueue(queue)
		if err != nil {
			log.Print(err.Error())
		}
		switch queue {
		case USERDEPOSIT, USERWIN:
			service.eventManager.ListenningToqueue(topic, service.AddBalance)
		case USERWITHDRAW, USERBET:
			service.eventManager.ListenningToqueue(topic, service.SubtractBalance)
		case USERREQUESTBET, USERREQUESTWITHDRAW:
			service.eventManager.ListenningToqueue(topic, service.CheckMoney)
		}

	}
}

func (service *userService) AddBalance(data []byte) error {
	service.lock.Lock()
	defer service.lock.Unlock()
	addmoney := domain.AddMoneyEvent{}
	err := json.Unmarshal(data, &addmoney)
	if err != nil {
		return err
	}
	err = service.repository.AddMoney(addmoney.UserId, addmoney.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) SubtractBalance(data []byte) error {
	service.lock.Lock()
	defer service.lock.Unlock()
	subtractmoney := domain.SubtractMoneyEvent{}
	err := json.Unmarshal(data, &subtractmoney)
	if err != nil {
		return err
	}
	err = service.repository.SubtractMoney(subtractmoney.UserId, subtractmoney.Amount)
	if err != nil {
		return err
	}
	service.UnReserveMoney(subtractmoney.UserId, subtractmoney.Amount, subtractmoney.Hash)
	return nil
}

func (service *userService) CheckMoney(data []byte) error {
	checkMoney := domain.CheckMoneyEvent{}
	err := json.Unmarshal(data, &checkMoney)
	if err != nil {
		return err
	}
	balance, err := service.repository.GetUserBalance(checkMoney.UserId)
	if err != nil {
		return err
	}
	confirmWithdraw := &domain.ConfirmMoneyEvent{
		Hash:        checkMoney.Hash,
		CanWithDraw: true,
	}
	reservedMoney, ok := reserveMoney[checkMoney.UserId]
	if ok {
		balance -= reservedMoney.Amount
	}
	if balance < checkMoney.Amount {
		confirmWithdraw.CanWithDraw = false
	}
	data, err = json.Marshal(confirmWithdraw)
	if err != nil {
		return err
	}
	err = service.eventManager.Publish(USERCONFIRMWITHDRAW, data)
	if err != nil {
		return err
	}
	if confirmWithdraw.CanWithDraw {
		service.ReserveMoney(checkMoney.UserId, checkMoney.Amount, checkMoney.Hash)
	}
	return nil
}

func (service *userService) ReserveMoney(userid string, money float64, hash string) {
	service.lock.Lock()
	defer service.lock.Unlock()
	reserver, ok := reserveMoney[userid]
	if !ok {
		reserveMoney[userid] = struct {
			Amount float64
			Hash   string
		}{
			Amount: money, Hash: hash,
		}
	}
	newReserver := struct {
		Amount float64
		Hash   string
	}{
		Amount: money + reserver.Amount, Hash: hash,
	}
	reserveMoney[userid] = newReserver
}

func (service *userService) UnReserveMoney(userid string, money float64, hash string) {
	service.lock.Lock()
	defer service.lock.Unlock()
	for _, reserver := range reserveMoney {
		if reserver.Hash == hash {
			reserver.Amount -= money
		}
	}
}

func hasFromPassword(password string) string {
	data, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(data)
}

