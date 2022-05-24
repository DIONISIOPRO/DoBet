package event

import (
	"encoding/json"

	"github.com/dionisiopro/dobet_payment/domain"
	"github.com/streadway/amqp"
)

type EventProcessor interface {
	CreateUser(user domain.User) error
	UpdateUser(userid string, user domain.User) error
	DeleteUser(userid string) error
	PayBet(userid, betId string, amount float64) error
	Deposit(domain.Deposit) error
}

type EventSubscriber interface {
	SubscribeToQueue(name string) <-chan amqp.Delivery
}

type Listenner interface {
	Listenning(done <-chan bool)
}
type RabbitMQListenner struct {
	Listenners []Listenner
}

func (l *RabbitMQListenner) Listenning(done <-chan bool) {
	for _, l := range l.Listenners {
		l.Listenning(done)
	}

}

func (l *RabbitMQListenner) AddListenner(listenner Listenner) {
	l.Listenners = append(l.Listenners, listenner)
}

type BetCreatedListenner struct {
	subscriber EventSubscriber
	Processor  EventProcessor
}

type UserCreatedListenner struct {
	subscriber EventSubscriber
	Processor  EventProcessor
}

type UserUpdatedListenner struct {
	subscriber EventSubscriber
	Processor  EventProcessor
}

type UserBetWinListenner struct {
	subscriber EventSubscriber
	Processor  EventProcessor
}

type UserDeletedListenner struct {
	subscriber EventSubscriber
	Processor  EventProcessor
}

func (l UserDeletedListenner) Listenning(done <-chan bool) {
	queue := l.subscriber.SubscribeToQueue(domain.USERDELETED)
	for {
		select {
		case <-done:
			break
		case data := <-queue:
			event := domain.UserDeletedEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.Processor.DeleteUser(event.UserId)
			if err == nil {
				data.Ack(false)
			}
		}
	}
}

func (l UserBetWinListenner) Listenning(done <-chan bool) {
	queue := l.subscriber.SubscribeToQueue(domain.BETWIN)
	for {
		select {
		case <-done:
			break
		case data := <-queue:
			event := domain.UserWinEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.Processor.Deposit(domain.Deposit{User_Id: event.User_id, Amount: event.Amount})
			if err == nil {
				data.Ack(false)
			}

		}
	}
}

func (l UserUpdatedListenner) Listenning(done <-chan bool) {
	queue := l.subscriber.SubscribeToQueue(domain.USERUPDATED)
	for {
		select {
		case <-done:
			break
		case data := <-queue:
			event := domain.UserUpdateEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.Processor.UpdateUser(event.UserId, event.User)
			if err == nil {
				data.Ack(false)
			}

		}
	}
}

func (l UserCreatedListenner) Listenning(done <-chan bool) {
	queue := l.subscriber.SubscribeToQueue(domain.USERCREATED)
	for {
		select {
		case <-done:
			break
		case data := <-queue:
			event := domain.UserCreatedEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.Processor.CreateUser(event.User)
			if err == nil {
				data.Ack(false)
			}

		}
	}
}

func (l BetCreatedListenner) Listenning(done <-chan bool) {
	queue := l.subscriber.SubscribeToQueue(domain.BETCREATED)
	for {
		select {
		case <-done:
			break
		case data := <-queue:
			event := domain.BetCreatedEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err != nil {
				continue
			}
			err = l.Processor.PayBet(event.User_id, event.Bet_Id, event.Amount)
			if err == nil {
				data.Ack(false)
			}

		}
	}
}
