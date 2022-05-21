package listenner

import (
	"encoding/json"
	"log"

	"github.com/dionisiopro/dobet-bet/domain/event"
)

type confirmPaymentEventListenner struct {
	service IncomingEventProcessorService
	subscriber eventSubscriber
}



func (l confirmPaymentEventListenner) Listenning(){
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

func NewconfirmPaymentEventListenner(service IncomingEventProcessorService, subscriber eventSubscriber) *confirmPaymentEventListenner{
	return &confirmPaymentEventListenner{
		service: service,
		subscriber: subscriber,
	}
}