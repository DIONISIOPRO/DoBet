package repositories

import "gitthub.com/dionisiopro/dobet/models"

func CreateBet(owner models.User, market models.Market, option models.BetOption, amount float64) error {
	return nil

}


func BetByUser(user_id string) []models.Bet{
	return []models.Bet{}
}

func BetByMatch(match_id string) []models.Bet{
	return []models.Bet{}
}

func Bets()[]models.Bet{
	return []models.Bet{}
}

func RunningBets() []models.Bet{
	return []models.Bet{}
}

func TotalRunningBetsMoney() float32{
	return 0
}