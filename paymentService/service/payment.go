package service

import (
	"errors"

	"gitthub.com/dionisiopro/dobet/repository"
)

type PaymentService interface {
	Deposit(amount float64, userid string) error
	Withdraw(amount float64, userid string) error
}

type paymentService struct {
	repository repository.PaymentRepository
}

func NewPaymentService(repository repository.PaymentRepository) PaymentService {
	return &paymentService{
		repository: repository,
	}
}

func (service *paymentService) Deposit(amount float64, userid string) error {
	if amount < 5 {

		return errors.New("pleasse deposit amount above or equal than 5")
	}
	if userid != "" {
		return errors.New("pleasse provide a valid user id")
	}
	
	return service.repository.Deposit(amount, userid)
}

func (service *paymentService) Withdraw(amount float64, userid string) error {
	if amount < 5 {

		return errors.New("pleasse withdraw amount above or equal than 5")
	}
	if userid != "" {
		return errors.New("pleasse provide a valid user id")
	}
	return service.repository.Withdraw(amount, userid)
}
