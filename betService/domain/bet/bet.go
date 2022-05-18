package bet

import (
	"github.com/dionisiopro/dobet-bet/domain/market"
	"github.com/go-playground/validator"
)

type BetBaseImpl struct {
	Bet_id     string          `json:"bet_id" bson:"bet_id"`
	Bet_owner  string          `json:"bet_owner" bson:"bet_owner" validate:"required"`
	BetGroup   []SingleBetImpl `json:"betgroup" validate:"required"`
	Status     string          `json:"status"`
	IsFinished bool            `json:"is_finished"`
}

type SingleBetImpl struct {
	League_id   string                 `json:"league_id" validate:"requied"`
	Match_id    string                 `json:"match_id" validate:"required"`
	IsProcessed bool                   `json:"isprocessed"`
	Amount      float64                `json:"totalamount" validate:"required"`
	Market      market.BetMarket   `json:"market" validate:"required"`
	Result      market.MatchResult `json:"result"`
	IsLose      bool                   `json:"is_lose"`
}

func (b BetBaseImpl) IsValid() bool {
	validate := validator.New()
	err := validate.Struct(b)
	if err != nil {
		return false
	}
	return true
}

func (b BetBaseImpl) GetGlobalOdd() float64 {
	odd := float64(0)
	for _, bet := range b.BetGroup {
		odd += bet.Market.GetSelectedOdd()
	}
	return odd
}

func (b BetBaseImpl) GetTotalAmount() float64 {
	amount := float64(0)
	for _, bet := range b.BetGroup {
		amount += bet.Amount
	}
	return amount
}

func (b BetBaseImpl) GetPotenctialWin() float64 {
	return b.GetGlobalOdd() * b.GetTotalAmount()

}

func (b BetBaseImpl) GetIsFinished() bool {
	for _, bet := range b.BetGroup {
		if !bet.IsProcessed {
			return false
		}
		continue
	}
	return true
}

func (b BetBaseImpl) IsLose() bool {
	loseCount := 0
	for _, bet := range b.BetGroup {
		if bet.IsLose {
			loseCount++
		}
		continue
	}
	return loseCount > 0
}

func (b *SingleBetImpl) SetResult(result interfaces.MatchResult) {
	b.Result = result
}

func (b SingleBetImpl) GetIsLose() bool {
	return b.Market.IsLose(b.Result)
}
