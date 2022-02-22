package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func CreateBet(owner models.User, market models.Market, option models.BetOption, amount float64) error {
	return repositories.CreateBet(owner, market, option,amount)

}

func BetByUser(user_id string) []models.Bet{
	return repositories.BetByUser(user_id)
}

func BetByMatch(match_id string) []models.Bet{
	return repositories.BetByUser(match_id)
}

func Bets()[]models.Bet{
	return repositories.Bets()
}

func RunningBets() []models.Bet{
	return repositories.RunningBets()
}

func TotalRunningBetsMoney() float32{
	return repositories.TotalRunningBetsMoney()
}