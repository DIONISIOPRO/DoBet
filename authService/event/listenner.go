package event

import (
	"encoding/json"
	"log"

	"github.com/dionisiopro/dobet-auth/domain"
)

type Listenner interface {
	Listenning(done chan bool)
}

type AuthEventListnner struct {
	Listenners []Listenner
}

func (a *AuthEventListnner) Listenning(done chan bool) {
	for _, listenner := range a.Listenners {
		go listenner.Listenning(done)
	}
}

func (a *AuthEventListnner) AddListenner(listenner Listenner) {
	a.Listenners = append(a.Listenners, listenner)
}

type CreateUserEventProcessor interface {
	CreateUser(user domain.User) error
}
type DeleteUserEventProcessor interface {
	DeleteUser(userId string) error
}
type UpdateUserEventProcessor interface {
	UpdateUser(userid string, user domain.User) error
}
type UserDeletedListenner struct {
	Subscriber EventSubscriber
	service    DeleteUserEventProcessor
}

type UserUpdatedListenner struct {
	Subscriber EventSubscriber
	service    UpdateUserEventProcessor
}

type UserCreatedListenner struct {
	Subscriber EventSubscriber
	service    CreateUserEventProcessor
}

func (l UserDeletedListenner) Listenning(done chan bool) {
	deliver, err := l.Subscriber.SubscribeToQueue(domain.USERDELETE)
	if err != nil {
		log.Print("error subscribing queue to listenning user deleted event")
	}
	for {
		select {
		case data := <-deliver:
			event := domain.DeleteUserEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.service.DeleteUser(event.UserId)
			if err != nil {
				data.Ack(false)
			}
		case <-done:
			break
		}
	}

}

func (l UserUpdatedListenner) Listenning(done chan bool) {
	deliver, err := l.Subscriber.SubscribeToQueue(domain.USERUPDATE)
	if err != nil {
		log.Print("error subscribing queue to listenning user deleted event")
	}
	for {
		select {
		case data := <-deliver:
			event := domain.UpdateUserEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.service.UpdateUser(event.UserId, event.User)
			if err != nil {
				data.Ack(false)
			}
		case <-done:
			break
		}
	}
}

func (l UserCreatedListenner) Listenning(done chan bool) {
	deliver, err := l.Subscriber.SubscribeToQueue(domain.USERCREATED)
	if err != nil {
		log.Print("error subscribing queue to listenning user deleted event")
	}
	for {
		select {
		case data := <-deliver:
			event := domain.AddUserEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.service.CreateUser(event.User)
			if err != nil {
				data.Ack(false)
			}
		case <-done:
			break
		}
	}
}
