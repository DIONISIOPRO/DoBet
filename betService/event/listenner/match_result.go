package listenner

import (
	"encoding/json"
	"log"

	"github.com/dionisiopro/dobet-bet/domain/event"
)
type matchResultEventListenner struct {
	subscriber eventSubscriber
	service IncomingEventProcessorService
}


func (l matchResultEventListenner) Listenning(done <-chan bool){
	queue, err:= l.subscriber.SubscribeToQueue(event.BetMatchResult)
	if err != nil{
		log.Println("error subscribing:", err.Error())
	}
	for {
		select {
		case <-done:
			break
		case data := <-queue:
			event := event.BetMatchConfirmEvent{}
			err := json.Unmarshal(data.Body, &event)
			if err == nil {
				continue
			}
			err = l.service.ConfirmBet(event.Bet_id)
			if err == nil {
				data.Ack(false)
			}
		}
	}
}

func NewmatchResultEventListenner(service IncomingEventProcessorService, subscriber eventSubscriber) *matchResultEventListenner{
	return &matchResultEventListenner{
		service: service,
		subscriber: subscriber,
	}
}