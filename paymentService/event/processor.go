package event

import (
	"encoding/json"
	"errors"

	"github.com/dionisiopro/dobet_payment/domain"
)

type PaymentRepository interface {
	Deposit(domain.Deposit) error
	Withdraw(domain.WithDraw) error
	CreateUser(user domain.User) error
	UpdateUser(user_id string, user domain.User) error
	DeleteUser(user_id string) error
}
type EventPublisher interface {
	Publish(queue string, data []byte) error
}
type RabbitEventProcessor struct {
	publisher   EventPublisher
	reposiotory PaymentRepository
}

func (p *RabbitEventProcessor) CreateUser(data []byte) error {
	event := domain.UserCreatedEvent{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return err
	}
	valid := event.IsValid()
	if !valid {
		return errors.New("invalid user")
	}
	err = p.reposiotory.CreateUser(event.User)
	if err != nil {
		return err
	}
	return nil
}

func (p *RabbitEventProcessor) UpdateUser(data []byte) error {
	event := domain.UserUpdateEvent{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return err
	}
	valid := event.IsValid()
	if !valid {
		return errors.New("invalid user")
	}
	err = p.reposiotory.UpdateUser(event.UserId, event.User)
	if err != nil {
		return err
	}
	return nil
}

func (p *RabbitEventProcessor) DeleteUser(data []byte) error {
	event := domain.UserDeletedEvent{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return err
	}
	err = p.reposiotory.DeleteUser(event.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (p *RabbitEventProcessor) Pay(data []byte) error {
	event := domain.BetCreatedEvent{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return err
	}
	witdraw := domain.WithDraw{}
	witdraw.Amount = event.Amount
	witdraw.Phone_number = event.Phone_number
	err = p.reposiotory.Withdraw(witdraw)
	if err != nil {
		return err
	}
	eventToPublish := domain.BetPayedEvent{Bet_id: event.Bet_Id}
	data, err = json.Marshal(eventToPublish)
	if err != nil {
		return err
	}
	return p.publisher.Publish(domain.USERBETPAYED, data)
}

func (p *RabbitEventProcessor) Refund(data []byte) error {
	event := domain.BetRefundEvent{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return err
	}
	deposit := domain.Deposit{}
	deposit.Amount = event.Amount
	deposit.Phone_number = event.Phone_number
	err = p.reposiotory.Deposit(deposit)
	if err != nil {
		return err
	}
	return nil
}
