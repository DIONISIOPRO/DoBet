package domain

type BetBaseImpl struct {
	Bet_id        string             `json:"bet_id" bson:"bet_id"`
	Bet_owner     string             `json:"bet_owner" bson:"bet_owner" validate:"required"`
	BetGroup      []Bet        `json:"betgroup" validate:"required"`
}

type SingleBetImpl struct {
	League_id   string `json:"league_id" validate:"requied"`
	Match_id    string    `json:"match_id" validate:"required"`
	IsProcessed bool      `json:"isprocessed"`
	Amount   float64            `json:"totalamount" validate:"required"`
	Market      BetMarket    `json:"market" validate:"required"`
}

func (b BetBaseImpl) GetGlobalOdd() float64{
	odd := float64(0)
	for index, bet := range b.BetGroup{
     odd += bet.Market.GetGlobalOdd()
	}
	return odd

}

func (b BetBaseImpl) GetTotalAmount() float64{
	amount := float64(0)
	for index, bet := range b.BetGroup{
     odd += bet.Amount
	}
	return amount
}

func (b BetBaseImpl ) GetPotenctialWin() float64{
	return b.GetGlobalOdd * b.GetTotalAmount

}

func (b BetBaseImpl) IsFinished() bool{
	for index, bet := range b.BetGroup{
		if !bet.IsProcessed{
			return false
		}
		continue
	   }
	return true
}

func (b BetBaseImpl) IsLose() bool{
	loseCount := 0
	for index, bet := range b.BetGroup{
		if bet.IsLose(){
			return true
		}
		continue
	   }
	return false
}

func (b SingleBetImpl) IsLose() bool{
	return b.Market.IsLose()
}

func (b Bet) IsValid() bool{

}





