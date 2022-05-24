package service

import (
	"errors"

	"github.com/dionisiopro/dobet-user/domain"
)

type (
	UserRepository interface {
		CreateUser(userRequest domain.User) (string, error)
		GetUserById(userId string) (domain.User, error)
		GetUserByPhone(phone string) (domain.User, error)
		GetUsers(startIndex, perpage int64) ([]domain.User, error)
		DeleteUser(userid string) error
		UpdateUser(userid string, user domain.User) error
		GetUserBalance(userId string) (float64, error)
		AddMoney(userId string, amount float64) error
		SubtractMoney(userId string, amount float64) error
	}
	UserService struct {
		repository         UserRepository
		userEventPublisher EventPublisher
	}
	EventPublisher interface {
		Publish(name string, event domain.Event) error
	}
)

func NewUserService(
	userRepository UserRepository, eventPublisher EventPublisher) *UserService {
	userService := &UserService{
		repository:         userRepository,
		userEventPublisher: eventPublisher,
	}
	return userService
}

func (service *UserService) CreateUser(user domain.User) (string, error) {
	err := user.Validate()
	if err != nil {
		return "", err
	}
	user.Account_balance = 0
	user.IsAdmin = false
	name, err := service.repository.CreateUser(user)
	if err != nil {
		return "", err
	}
	userCreated := &domain.UserCreatedEvent{
		User: user,
	}
	err = service.userEventPublisher.Publish(domain.USERCREATED, userCreated)
	return name, err
}

func (service *UserService) GetUsers(page, perpage int64) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	tmpusers, err := service.repository.GetUsers(startIndex, perpage)
	var users []domain.User

	if err != nil {
		return nil, err
	}
	for _, _user := range tmpusers {
		balance, err := service.getBalanceAndAtachToUser(_user.User_id)
		if err != nil {
			return nil, err
		}
		_user.Account_balance = balance
		users = append(users, _user)
	}
	return users, nil
}

func (service *UserService) GetUserById(userId string) (domain.User, error) {
	if userId == "" {
		return domain.User{}, errors.New("user id is empty")
	}
	user, err := service.repository.GetUserById(userId)
	if err != nil {
		return domain.User{}, err
	}
	balance, err := service.getBalanceAndAtachToUser(user.User_id)
	user.Account_balance = balance
	return user, nil
}

func (service *UserService) GetUserByPhone(phone string) (domain.User, error) {
	if phone == "" {
		return domain.User{}, errors.New("user id is empty")
	}
	user, err := service.repository.GetUserByPhone(phone)
	if err != nil {
		return domain.User{}, err
	}
	balance, err := service.getBalanceAndAtachToUser(user.User_id)
	user.Account_balance = balance
	return user, nil
}

func (service *UserService) DeleteUser(userid string) error {
	if userid == "" {
		return errors.New("user id is empty")
	}
	err := service.repository.DeleteUser(userid)
	if err != nil {
		return err
	}
	userDeletedEvent := domain.UserDeletedEvent{
		UserId: userid,
	}
	err = service.userEventPublisher.Publish(domain.USERDELETE, userDeletedEvent)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUser(userid string, user domain.User) error {
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
	return service.userEventPublisher.Publish(domain.USERUPDATE, userUpdateEvent)
}

func (service *UserService) getBalanceAndAtachToUser(userid string) (float64, error) {
	// make a call to payment service and get the balance of user
	// atach the balane to the curent user
	return 0.0, nil
}
