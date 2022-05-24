package event

import "github.com/streadway/amqp"

func NewEventSubscriber(conn *amqp.Connection) *RabbitMQEventSubscriber {
	return &RabbitMQEventSubscriber{
		Conn: conn,
	}
}
func NewEventListenner() *RabbitMQListenner {
	return &RabbitMQListenner{
		Listenners: []Listenner{},
	}
}

func NewRabbitMQEventPublisher(conn *amqp.Connection) *RabbiMQEventPublisher {
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return &RabbiMQEventPublisher{
		channel: channel,
	}
}

func NewBetCreatedListenner(subscriber EventSubscriber, Processor EventProcessor) *BetCreatedListenner {
	return &BetCreatedListenner{
		subscriber: subscriber,
		Processor:  Processor,
	}
}

func NewUserDeletedListenner(subscriber EventSubscriber, Processor EventProcessor) *UserDeletedListenner {
	return &UserDeletedListenner{
		subscriber: subscriber,
		Processor:  Processor,
	}
}

func NewUserBetWinListenner(subscriber EventSubscriber, Processor EventProcessor) *UserBetWinListenner {
	return &UserBetWinListenner{
		subscriber: subscriber,
		Processor:  Processor,
	}
}

func NewUserUpdatedListenner(subscriber EventSubscriber, Processor EventProcessor) *UserUpdatedListenner {
	return &UserUpdatedListenner{
		subscriber: subscriber,
		Processor:  Processor,
	}
}

func NewUserCreatedListenner(subscriber EventSubscriber, Processor EventProcessor) *UserCreatedListenner {
	return &UserCreatedListenner{
		subscriber: subscriber,
		Processor:  Processor,
	}
}
