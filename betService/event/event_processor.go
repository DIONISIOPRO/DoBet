package event

type IncomingEventProcessorRepository interface {
	ConfirmBet(bet_id string) error
	ActiveBet(bet_id string) error
	CancelBet(bet_id string) error
}

type IncomingEventBetProcessor interface {
	ProcessMatchResultInBet([]byte)
}

type IncomingEventProcessor struct {
	repository  IncomingEventProcessorRepository
	betPocessor IncomingEventBetProcessor
}

func NewIncomingEventProcessor(repository IncomingEventProcessorRepository) IncomingEventProcessor {
	incomingEventHandler := IncomingEventProcessor{
		repository: repository,
	}
	return incomingEventHandler
}

func (h IncomingEventProcessor) ConfirmBet(bet_id string) error {
	return h.repository.ConfirmBet(bet_id)
}

func (h IncomingEventProcessor) ActiveBet(bet_id string) error {
	return h.repository.ActiveBet(bet_id)
}

func (h IncomingEventProcessor) ProcessMatchResultInBet(data []byte) {
	h.betPocessor.ProcessMatchResultInBet(data)
}

func (h IncomingEventProcessor)CancelBet(bet_id string) error{
	return h.CancelBet(bet_id)
}
