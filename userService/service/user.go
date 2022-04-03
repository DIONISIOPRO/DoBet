package service

import (
	"encoding/json"
	"errors"
	"github/namuethopro/dobet-user/domain"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

const (
	USERDELETE          = "user.delete"
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
	USERLOGIN, USERLOGOUT, USERSIGNUP, USERDELETE, USERUPDATE, USERCONFIRMWITHDRAW, USERCONFIRMBET,
}

type (
	userRepository interface {
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
		Publish(name string, body []byte) error
		ListenningToqueue(queue <-chan amqp.Delivery, f func([]byte))
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

func (service *userService) GetUsers(page, perpage int64) (domain.UsersResponse, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	users, err := service.repository.GetUsers(startIndex, perpage)

	if err != nil {
		return domain.UsersResponse{}, err
	}
	userResponse := domain.UsersResponse{}
	for _, user := range users {
		userResponse.Users = append(userResponse.Users, user.ToUserResponse())
	}
	return userResponse, nil
}

func (service *userService) GetUserById(userId string) (domain.UserResponse, error) {
	user, err := service.repository.GetUserById(userId)
	if err != nil {
		return domain.UserResponse{}, err
	}
	return user.ToUserResponse(), nil
}

func (service *userService) GetUserByPhone(phone string) (domain.UserResponse, error) {
	user, err := service.repository.GetUserByPhone(phone)
	if err != nil {
		return domain.UserResponse{}, err
	}
	return user.ToUserResponse(), nil
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
	data, err := json.Marshal(userDeletedEvent)
	if err != nil {
		return err
	}
	return service.eventManager.Publish(USERDELETE, data)
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
	data, err := json.Marshal(userUpdateEvent)
	if err != nil {
		return err
	}
	return service.eventManager.Publish(USERUPDATE, data)
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

func (service *userService) AddBalance(data []byte) {
	service.lock.Lock()
	defer service.lock.Unlock()
	addmoney := domain.AddMoneyEvent{}
	err := json.Unmarshal(data, &addmoney)
	if err != nil {
		log.Fatal(err)
	}
	err = service.repository.AddMoney(addmoney.UserId, addmoney.Amount)
	if err != nil {
		log.Fatal(err)
	}
}

func (service *userService) SubtractBalance(data []byte) {
	service.lock.Lock()
	defer service.lock.Unlock()
	subtractmoney := domain.SubtractMoneyEvent{}
	err := json.Unmarshal(data, &subtractmoney)
	if err != nil {
		log.Fatal(err)
	}
	err = service.repository.SubtractMoney(subtractmoney.UserId, subtractmoney.Amount)
	if err != nil {
		log.Fatal(err)
	}
	service.UnReserveMoney(subtractmoney.UserId, subtractmoney.Amount, subtractmoney.Hash)
}

func (service *userService) CheckMoney(data []byte) {
	checkMoney := domain.CheckMoneyEvent{}
	err := json.Unmarshal(data, &checkMoney)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := service.repository.GetUserBalance(checkMoney.UserId)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	err = service.eventManager.Publish(USERCONFIRMWITHDRAW, data)
	if err != nil {
		log.Fatal(err)
	}
	if confirmWithdraw.CanWithDraw {
		service.ReserveMoney(checkMoney.UserId, checkMoney.Amount, checkMoney.Hash)
	}

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
