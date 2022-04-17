package event

import (
	"encoding/json"

	"github/namuethopro/dobet-auth/domain"
)

type IncomingEventProcessorRepository interface {
	AddUser(user domain.User) error
	UpdateUser(userid string, user domain.User) error
	RemoveUser(userId string) error
}

type IncomingEventProcessor struct {
	repository IncomingEventProcessorRepository
}

func NewIncomingEventProcessor(repository IncomingEventProcessorRepository) IncomingEventProcessor {
	incomingEventHandler := IncomingEventProcessor{
		repository: repository,
	}
	return incomingEventHandler
}

func (eHandler IncomingEventProcessor) AddUser(data []byte) error {
	var AddUser = domain.AddUserEvent{}
	err := json.Unmarshal(data, &AddUser)
	if err != nil {
		return err
	}
	err = eHandler.repository.AddUser(AddUser.User)
	if err != nil {
		return err
	}
	return nil
}

func (eHandler IncomingEventProcessor) UpdateUser(data []byte) error {
	var user = domain.UpdateUserEvent{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		return err
	}
	err = eHandler.repository.UpdateUser(user.UserId, user.User)
	if err != nil {
		return err
	}
	return nil
}

func (eHandler IncomingEventProcessor) RemoveUser(data []byte) error {
	deleteEvent := domain.DeleteUserEvent{}
	err := json.Unmarshal(data, &deleteEvent)
	if err != nil {
		return err
	}
	err = eHandler.RemoveUser([]byte(deleteEvent.UserId))
	if err != nil {
		return err
	}
	return nil
}
