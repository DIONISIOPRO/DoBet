package service

import (
	"encoding/json"
	"errors"

	"github.com/dionisiopro/dobet_payment/domain"
)

var (
	phonenumberinvaliderr = errors.New("pleasse provide a valid user phone number")
	amountlessthan5Err    = errors.New("pleasse withdraw amount above or equal than 5")
)

type PaymentRepository interface {
	Deposit(domain.Deposit) error
	Withdraw(domain.WithDraw) error
	CreateUser(user domain.User) error
	DeleteUser(user_id string) error
}
type PaymentApi interface {
	Deposit(phone_number string, amount float64) error
	Withdraw(phone_number string, amount float64) error
}
type EventPublisher interface {
	Publish(queue string, data []byte) error
}
type paymentService struct {
	repository PaymentRepository
	Api        PaymentApi
	publisher  EventPublisher
}

func NewPaymentService(repository PaymentRepository, Api PaymentApi) *paymentService {
	return &paymentService{
		repository: repository,
		Api:        Api,
	}
}

func (s *paymentService) Deposit(deposit domain.Deposit) error {
	if deposit.Amount < 5 {

		return amountlessthan5Err
	}
	if deposit.Phone_number != "" {
		return phonenumberinvaliderr
	}
	data, err := prepareBalanceEvent(deposit.Amount, deposit.Phone_number)
	if err != nil {
		return err
	}
	err = s.repository.Deposit(deposit)
	if err != nil {
		return err
	}
	return s.publisher.Publish(domain.USERBALANCEUPDATED, data)
}

func (s *paymentService) WithDraw(withdraw domain.WithDraw) error {
	if withdraw.Amount < 5 {

		return amountlessthan5Err
	}
	if withdraw.Phone_number != "" {
		return phonenumberinvaliderr
	}
	data, err := prepareBalanceEvent(withdraw.Amount, withdraw.Phone_number)
	if err != nil {
		return err
	}
	err = s.repository.Withdraw(withdraw)
	return s.publisher.Publish(domain.USERBALANCEUPDATED, data)
}

func prepareBalanceEvent(amount float64, phone_number string) ([]byte, error) {
	balanceevent := domain.BalanceUpdateEvent{}
	balanceevent.Amount = amount
	balanceevent.Phone_number = phone_number
	data, err := json.Marshal(balanceevent)
	if err != nil {
		return nil, err
	}
	return data, nil
}
