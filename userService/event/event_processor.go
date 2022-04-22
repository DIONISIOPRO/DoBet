package event

import (
	"encoding/json"
	"sync"

	"github.com/namuethopro/dobet-user/domain"
)

type IncomingEventProcessorRepository interface {
	GetUserBalance(userId string) (float64, error)
	AddMoney(userId string, amount float64) error
	SubtractMoney(userId string, amount float64) error
}
type IncomingEventProcessorConfirmEventManager interface {
	Publish(name string, event domain.Event) error
}
type IncomingEventProcessorReserveMoneyManager interface {
	ReserveMoney(userid string, money float64, hash string)
	UnReserveMoney(userid string, money float64, hash string)
	GetReservedMoneyByUserId(userId string) float64
}

type IncomingEventProcessor struct {
	lock          *sync.Mutex
	repository    IncomingEventProcessorRepository
	eventManager  IncomingEventProcessorConfirmEventManager
	moneyReserver IncomingEventProcessorReserveMoneyManager
}

func NewIncomingEventProcessor(lock *sync.Mutex, repository IncomingEventProcessorRepository, eventManger IncomingEventProcessorConfirmEventManager, moneyReserver IncomingEventProcessorReserveMoneyManager) IncomingEventProcessor {
	incomingEventHandler := IncomingEventProcessor{
		lock:          lock,
		repository:    repository,
		eventManager:  eventManger,
		moneyReserver: moneyReserver,
	}
	return incomingEventHandler
}

func (eHandler IncomingEventProcessor) AddBalance(data []byte) error {
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

func (eHandler IncomingEventProcessor) SubtractBalance(data []byte) error {
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
	eHandler.unreserveMoney(subtractmoney.UserId, subtractmoney.Amount, subtractmoney.Hash)
	return nil
}

func (eHandler IncomingEventProcessor) CheckMoney(data []byte) error {
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
	reservedMoney := eHandler.moneyReserver.GetReservedMoneyByUserId(checkMoney.UserId)
	if balance - reservedMoney < checkMoney.Amount {
		confirmWithdraw.CanWithDraw = false
	}
	err = eHandler.confirmMoney(domain.USERCONFIRMWITHDRAW, confirmWithdraw)
	if err != nil {
		return err
	}
	if confirmWithdraw.CanWithDraw {
		eHandler.reserveMoney(checkMoney.UserId, checkMoney.Amount, checkMoney.Hash)
	}
	return nil
}

func (eHandler IncomingEventProcessor) reserveMoney(userId string, money float64, hash string) {
	eHandler.moneyReserver.ReserveMoney(userId, money, hash)
}
func (eHandler IncomingEventProcessor) unreserveMoney(userId string, money float64, hash string) {
	eHandler.moneyReserver.UnReserveMoney(userId, money, hash)
}

func (eHandler IncomingEventProcessor) confirmMoney(queue string, event domain.Event) error {
	return eHandler.eventManager.Publish(queue, event)
}
