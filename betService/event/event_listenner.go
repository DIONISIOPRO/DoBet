package event

import (
	"encoding/json"
	"log"

	"github.com/dionisiopro/dobet-bet/domain/event"
	"github.com/dionisiopro/dobet-bet/domain/result"
	"github.com/streadway/amqp"
)
type IncomingEventProcessorService interface {
	ConfirmBet(bet_id string) error
	ActiveBet(bet_id string) error
	CancelBet(bet_id string) error
	ProcessMatchResultInBet(result.MatchResult) error
}
type eventSubscriber interface{
	SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
}

type ConfirmPaymentEventListenner struct {
	service IncomingEventProcessorService
	subscriber eventSubscriber
}

type ConfirmMatchEventListenner struct {
	service IncomingEventProcessorService
	subscriber eventSubscriber
}

type MatchResultEventListenner struct {
	subscriber eventSubscriber
	service IncomingEventProcessorService
}


func (l MatchResultEventListenner) Listenning(){
	queue, err:= l.subscriber.SubscribeToQueue(event.BetMatchResult)
	if err != nil{
		log.Println("error subscribing:", err.Error())
	}
	for data := range queue{
		event := event.BetResultMatchEvent{}
		err := json.Unmarshal(data.Body,&event)
		if err == nil{
			continue
		}
		err = l.service.ProcessMatchResultInBet(event.Result)
		if err == nil{
			data.Ack(false)
		}
	}
}

func (l ConfirmMatchEventListenner) Listenning(){
	queue, err:= l.subscriber.SubscribeToQueue(event.BetMatchConfirm)
	if err != nil{
		log.Println("error subscribing:", err.Error())
	}
	for data := range queue{
		event := event.BetMatchConfirmEvent{}
		err := json.Unmarshal(data.Body,&event)
		if err == nil{
			continue
		}
		err = l.service.ConfirmBet(event.Bet_id)
		if err == nil{
			data.Ack(false)
		}
	}
}

func (l ConfirmPaymentEventListenner) Listenning(){
	queue, err:= l.subscriber.SubscribeToQueue(event.BetPaymentConfirm)
	if err != nil{
		log.Println("error subscribing:", err.Error())
	}
	for data := range queue{
		event := event.BetPaymentConfirmEvent{}
		err := json.Unmarshal(data.Body,&event)
		if err == nil{
			continue
		}
		err = l.service.ActiveBet(event.Bet_id)
		if err == nil{
			data.Ack(false)
		}
	}

}