package bet

import (
	"github.com/dionisiopro/dobet-bet/domain/market"
	"github.com/dionisiopro/dobet-bet/domain/result"
	"github.com/go-playground/validator"
)
const(
	Active = "active"
	Canceled = "canceled"
	Confirmed = "confirmed"
)
type BetBase struct {
	Bet_id     string          `json:"bet_id" bson:"bet_id"`
	Bet_owner  string          `json:"bet_owner" bson:"bet_owner" validate:"required"`
	BetGroup   []SingleBetImpl `json:"betgroup" validate:"required"`
	Status     string          `json:"status"`
	IsFinished bool            `json:"is_finished"`
}

type SingleBetImpl struct {
	League_id   string             `json:"league_id" validate:"requied"`
	Match_id    string             `json:"match_id" validate:"required"`
	IsProcessed bool               `json:"isprocessed"`
	Amount      float64            `json:"totalamount" validate:"required"`
	Market      market.MatchMarket `json:"market" validate:"required"`
	Result      result.MatchResult `json:"result"`
	IsLose      bool               `json:"is_lose"`
}

func (b BetBase) IsValid() bool {
	validate := validator.New()
	err := validate.Struct(b)
	if err != nil {
		return false
	}
	return true
}

func (b BetBase) GetGlobalOdd() float64 {
	odd := float64(0)
	for _, bet := range b.BetGroup {
		odd += bet.Market.GetSelectedOdd()
	}
	return odd
}

func (b BetBase) GetTotalAmount() float64 {
	amount := float64(0)
	for _, bet := range b.BetGroup {
		amount += bet.Amount
	}
	return amount
}

func (b BetBase) GetPotenctialWin() float64 {
	return b.GetGlobalOdd() * b.GetTotalAmount()

}

func (b BetBase) GetIsFinished() bool {
	for _, bet := range b.BetGroup {
		if !bet.IsProcessed {
			return false
		}
		continue
	}
	return true
}

func (b BetBase) IsLose() bool {
	loseCount := 0
	for _, bet := range b.BetGroup {
		if bet.IsLose {
			loseCount++
		}
		continue
	}
	return loseCount > 0
}

func (b *SingleBetImpl) SetResult(result result.MatchResult) {
	b.Result = result
}

func (b SingleBetImpl) GetIsLose() bool {
	return b.Market.IsLose(b.Result)
}
