package event

import (
	"encoding/json"

	"github.com/dionisiopro/dobet-bet/domain/market"
	"github.com/dionisiopro/dobet-bet/domain/result"
)

const (
	BetCreated        = ""
	BetMatchConfirm   = ""
	BetCanceled       = ""
	BetPaymentConfirm = ""
	BetMatchResult = ""
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
		Market market.MatchMarket
	}

	BetResultMatchEvent struct {
		Result result.MatchResult
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

func (b BetCreatedEvent) ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetDepositEvent) ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetPaymentConfirmEvent) ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetMatchConfirmEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}

func (b BetCanceledEvent)ToByteArray() ([]byte, error){
	return json.Marshal(b)
}