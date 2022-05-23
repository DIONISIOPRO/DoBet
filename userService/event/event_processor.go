package event

import (
	"encoding/json"
	"github.com/dionisiopro/dobet-user/domain"
)

type IncomingEventProcessorRepository interface {
	AddMoney(userId string, amount float64) error
	SubtractMoney(userId string, amount float64) error
}

type IncomingEventProcessor struct {
	repository    IncomingEventProcessorRepository
}

func NewIncomingEventProcessor( repository IncomingEventProcessorRepository) IncomingEventProcessor {
	incomingEventHandler := IncomingEventProcessor{
		repository:    repository,
	}
	return incomingEventHandler
}

func (eHandler IncomingEventProcessor) AddBalance(data []byte) error {
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
	subtractmoney := domain.SubtractMoneyEvent{}
	err := json.Unmarshal(data, &subtractmoney)
	if err != nil {
		return err
	}
	err = eHandler.repository.SubtractMoney(subtractmoney.UserId, subtractmoney.Amount)
	if err != nil {
		return err
	}
	return nil
}

