package service

import (
	"github.com/namuethopro/dobet-user/domain"
	"sync"
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
	userService struct {
		repository         UserRepository
		eventProcessor     EventProcessor
		eventListenner     EventListenner
		userEventPublisher EventPublisher
		lock               *sync.Mutex
	}
	EventPublisher interface {
		Publish(name string, event domain.Event) error
	}
	EventListenner interface {
		ListenningToqueues(done <-chan bool)
	}
	EventProcessor interface {
		SubtractBalance(data []byte) error
		CheckMoney(data []byte) error
		AddBalance(data []byte) error
	}
)

func newUserService(
	userRepository UserRepository,
	eventPublisher EventPublisher,
	eventListenner EventListenner,
	eventProcessor EventProcessor,
	lock *sync.Mutex) userService {
	userService := userService{
		repository:         userRepository,
		userEventPublisher: eventPublisher,
		eventProcessor:     eventProcessor,
		eventListenner:     eventListenner,
		lock:               lock,
	}
	return userService
}

func (service userService) CreateUser(user domain.User) (string, error) {
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

func (service userService) GetUsers(page, perpage int64) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	users, err := service.repository.GetUsers(startIndex, perpage)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service userService) GetUserById(userId string) (domain.User, error) {
	if userId == "" {
		return domain.User{}, errors.New("user id is empty")
	}
	user, err := service.repository.GetUserById(userId)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (service userService) GetUserByPhone(phone string) (domain.User, error) {
	if phone == "" {
		return domain.User{}, errors.New("user id is empty")
	}
	user, err := service.repository.GetUserByPhone(phone)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (service userService) DeleteUser(userid string) error {
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

func (service userService) UpdateUser(userid string, user domain.User) error {
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

func (service *userService) StartListenningEvents(done <-chan bool) {
	service.eventListenner.ListenningToqueues(done)
}
