package listenner

import (
	"encoding/json"
	"log"

	"github.com/dionisiopro/dobet-bet/domain/event"
)

type confirmMatchEventListenner struct {
	service    IncomingEventProcessorService
	subscriber eventSubscriber
}

func (l confirmMatchEventListenner) Listenning(done <-chan bool) {
	queue, err := l.subscriber.SubscribeToQueue(event.BetMatchConfirm)
	if err != nil {
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

func NewConfirmMatchEventListenner(service IncomingEventProcessorService, subscriber eventSubscriber) *confirmMatchEventListenner {
	return &confirmMatchEventListenner{
		service:    service,
		subscriber: subscriber,
	}
}
