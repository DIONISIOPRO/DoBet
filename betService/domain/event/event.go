package event

import (
	"encoding/json"

	"github.com/dionisiopro/dobet-bet/domain/markets"
)

const (
	BetCreated        = ""
	BetMatchConfirm   = ""
	BetCanceled       = ""
	BetPaymentConfirm = ""
	BetDeposit        = ""
)

var (
	EventsToPublish    = []string{BetDeposit, BetCreated}
	EventsToListenning = []string{BetMatchConfirm, BetPaymentConfirm, BetCanceled}
)

type (
	BetCreatedEvent struct {
		User_id string
		Bet_id string
		Match_idS []string
		Market markets.MatchMarketBase
	}

	BetMatchConfirmEvent struct {
		Bet_id string
		Confirmed bool
	}

	BetCanceledEvent struct {
		Bet_id string
	}

	BetPaymentConfirmEvent struct {
		Bet_id string
	}

	BetDepositEvent struct {
		User_id string
		Amount float64
	}
)

func (b BetCreatedEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetDepositEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetPaymentConfirmEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetMatchConfirmEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetCanceledEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}