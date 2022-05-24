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
	UpdateUser(userid string, user domain.User) error
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

func NewPaymentService(repository PaymentRepository, publisher  EventPublisher, Api PaymentApi) *paymentService {
	return &paymentService{
		repository: repository,
		Api:        Api,
		publisher: publisher,
	}
}

func (s *paymentService) CreateUser(user domain.User) error {
	return s.repository.CreateUser(user)
}
func (s *paymentService) UpdateUser(userid string, user domain.User) error {
	return s.repository.UpdateUser(userid, user)
}
func (s *paymentService) DeleteUser(userid string) error {
	return s.repository.DeleteUser(userid)
}

func (s *paymentService) PayBet(userid, betId string, amount float64) error {
	event := domain.BetPayedEvent{
		Bet_id: betId,
	}
	data, err := json.Marshal(event)
	if err != nil{
		return err
	}
	err = s.repository.Withdraw(domain.WithDraw{User_id: userid, Amount: amount})
	if err != nil {
		return err
	}
	s.publisher.Publish(domain.BETPAYED, data)
	return nil
}

func (s *paymentService) Deposit(deposit domain.Deposit) error {
	if deposit.Amount < 5 {

		return amountlessthan5Err
	}
	if deposit.User_Id != "" {
		return phonenumberinvaliderr
	}
	err := s.repository.Deposit(deposit)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *paymentService) WithDraw(withdraw domain.WithDraw) error {
	if withdraw.Amount < 5 {
		return amountlessthan5Err
	}
	if withdraw.User_id != "" {
		return phonenumberinvaliderr
	}

	err := s.repository.Withdraw(withdraw)
	if err != nil {
		return err
	}
	return nil
}
