package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/namuethopro/dobet-user/domain"
	"github/namuethopro/dobet-user/event"
	"github/namuethopro/dobet-user/repository"
	"log"
	"sync"
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

var ReserveMoney = make(map[string]struct {
	Amount float64
	Hash   string
})
var QueuesToListenning = []string{
	USERREQUESTBET, USERREQUESTWITHDRAW, USERDEPOSIT, USERWIN, USERWITHDRAW,USERBET,
}
var QueuesToPublish = []string{
	USERLOGIN, USERLOGOUT, USERSIGNUP, USERDELETE, USERUPDATE, USERCONFIRMWITHDRAW, USERCONFIRMBET,
}


type UserService interface {
	Users(page, perpage int64) ([]domain.UserResponse, error)
	GetUserById(userId string) (domain.UserResponse, error)
	GetUserByPhone(phone string) (domain.UserResponse, error)
	DeleteUser(userid string) error
	UpdateUser(userid string, user domain.User) error
	ListenningToqueue()
}

type userService struct {
	repository   repository.UserRepository
	eventManager event.EventManager
	lock         *sync.Mutex
}

func NewUserService(userRepository repository.UserRepository, eventManager event.EventManager, lock *sync.Mutex) UserService {
	userService := &userService{
		repository:   userRepository,
		eventManager: eventManager,
		lock:         lock,
	}
	go func ()  {
		userService.eventManager.CreateQueues(QueuesToPublish)
		userService.eventManager.CreateQueues(QueuesToListenning)
	}()
	return userService
}

func (service *userService) Users(page, perpage int64) ([]domain.UserResponse, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	users, err := service.repository.Users(startIndex, perpage)

	if err != nil {
		return nil, err
	}
	userResponse := []domain.UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, user.ToResponse())
	}
	return userResponse, nil
}

func (service *userService) GetUserById(userId string) (domain.UserResponse, error) {
	user, err := service.repository.GetUserById(userId)
	if err != nil {
		return domain.UserResponse{}, err
	}
	fmt.Println("chegou aqui")
	return user.ToResponse(), nil
}

func (service *userService) GetUserByPhone(phone string) (domain.UserResponse, error) {
	user, err := service.repository.GetUserByPhone(phone)
	if err != nil {
		return domain.UserResponse{}, err
	}
	return user.ToResponse(), nil
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
	for _, queue := range QueuesToListenning {
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
	reservedMoney, ok := ReserveMoney[checkMoney.UserId]
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
	reserver, ok := ReserveMoney[userid]
	if !ok {
		ReserveMoney[userid] = struct {
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
	ReserveMoney[userid] = newReserver
}

func (service *userService) UnReserveMoney(userid string, money float64, hash string) {
	service.lock.Lock()
	defer service.lock.Unlock()
	for _, reserver := range ReserveMoney {
		if reserver.Hash == hash {
			reserver.Amount -= money
		}
	}
}
