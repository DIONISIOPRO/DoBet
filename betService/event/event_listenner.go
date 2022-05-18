package event

import "github.com/streadway/amqp"

type eventSubscriber interface{
	SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
}
type eventProcessor interface{
	Process(id interface{}) error
}

type ConfirmPaymentEventListenner struct {
	subscriber eventSubscriber
	processor eventProcessor
}

type ConfirmMatchEventListenner struct {
	subscriber eventSubscriber
	processor eventProcessor
}

type MatchResultEventListenner struct {
	subscriber eventSubscriber
	processor eventProcessor
}


func (l MatchResultEventListenner) Listenning(){
	

}

func (l ConfirmMatchEventListenner) Listenning(){

}

func (l ConfirmPaymentEventListenner) Listenning(){

}