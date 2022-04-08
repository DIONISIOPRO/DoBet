package service

import (
	"encoding/json"
	"sync"

	"github/namuethopro/dobet-user/domain"
)

type IventHandlerUserRepository interface {
	GetUserBalance(userId string) (float64, error)
	AddMoney(userId string, amount float64) error
	SubtractMoney(userId string, amount float64) error
}
type IventHandlerUserEventManager interface {
	Publish(name string, event domain.Event) error
}
type ReserveMoneyHandler interface {
	ReserveMoney(userid string, money float64, hash string)
	UnReserveMoney(userid string, money float64, hash string)
}
type IncomingEventHandler struct {
	lock          *sync.Mutex
	repository    IventHandlerUserRepository
	eventManager  IventHandlerUserEventManager
	moneyReserver ReserveMoneyHandler
}

func NewIncomingEventHandler(lock *sync.Mutex, repository IventHandlerUserRepository, eventManger IventHandlerUserEventManager, moneyReserver ReserveMoneyHandler) IncomingEventHandler {
	incomingEventHandler := IncomingEventHandler{
		lock:          lock,
		repository:    repository,
		eventManager:  eventManger,
		moneyReserver: moneyReserver,
	}
	return incomingEventHandler
}

func (eHandler IncomingEventHandler) AddBalance(data []byte) error {
	eHandler.lock.Lock()
	defer eHandler.lock.Unlock()
	var addmoney = domain.AddMoneyEvent{}
	err := json.Unmarshal(data, &addmoney)
	if err != nil {
		return err
	}
	err = eHandler.repository.AddMoney("id", 10)
	if err != nil {
		return err
	}
	return nil
}

func (eHandler IncomingEventHandler) SubtractBalance(data []byte) error {
	eHandler.lock.Lock()
	defer eHandler.lock.Unlock()
	subtractmoney := domain.SubtractMoneyEvent{}
	err := json.Unmarshal(data, &subtractmoney)
	if err != nil {
		return err
	}
	err = eHandler.repository.SubtractMoney(subtractmoney.UserId, subtractmoney.Amount)
	if err != nil {
		return err
	}
	eHandler.moneyReserver.UnReserveMoney(subtractmoney.UserId, subtractmoney.Amount, subtractmoney.Hash)
	return nil
}

func (eHandler IncomingEventHandler) CheckMoney(data []byte) error {
	checkMoney := domain.CheckMoneyEvent{}
	err := json.Unmarshal(data, &checkMoney)
	if err != nil {
		return err
	}
	balance, err := eHandler.repository.GetUserBalance(checkMoney.UserId)
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

	err = eHandler.eventManager.Publish(USERCONFIRMWITHDRAW, confirmWithdraw)
	if err != nil {
		return err
	}
	if confirmWithdraw.CanWithDraw {
		eHandler.moneyReserver.ReserveMoney(checkMoney.UserId, checkMoney.Amount, checkMoney.Hash)
	}
	return nil
}
