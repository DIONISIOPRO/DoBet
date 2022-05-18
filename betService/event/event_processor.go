package event

import "github.com/dionisiopro/dobet-bet/domain/result"


type IncomingEventProcessorRepository interface {
	ConfirmBet(bet_id string) error
	ActiveBet(bet_id string) error
	CancelBet(bet_id string) error
}

type IncomingEventBetProcessor interface {
	ProcessMatchResultInBet(result.MatchResult) error
}


type ConfirmPaymentEventProcessor struct {
	repository  IncomingEventProcessorRepository
}

type ConfirmMatchEventProcessor struct {
	repository  IncomingEventProcessorRepository
}

type MatchResultEventProcessor struct {
	betPocessor IncomingEventBetProcessor
}

func (p ConfirmPaymentEventProcessor) Process(id string) error {
	return p.repository.ActiveBet(id)
}

func (p ConfirmMatchEventProcessor) Process(id string) error {
	return p.repository.ConfirmBet(id)
}

func (p MatchResultEventProcessor) Process(result result.MatchResult) error {
	return p.betPocessor.ProcessMatchResultInBet(result)
}