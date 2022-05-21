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


func (l matchResultEventListenner) Listenning(){
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

func NewmatchResultEventListenner(service IncomingEventProcessorService, subscriber eventSubscriber) *matchResultEventListenner{
	return &matchResultEventListenner{
		service: service,
		subscriber: subscriber,
	}
}