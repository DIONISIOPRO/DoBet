package listenner

import (
	"github.com/dionisiopro/dobet-bet/domain/result"
	"github.com/streadway/amqp"
)

type IncomingEventProcessorService interface {
	ConfirmBet(bet_id string) error
	ActiveBet(bet_id string) error
	CancelBet(bet_id string) error
	ProcessMatchResultInBet(result.MatchResult) error
}
type eventSubscriber interface {
	SubscribeToQueue(name string) (<-chan amqp.Delivery, error)
}
