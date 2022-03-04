package helpers

import "gitthub.com/dionisiopro/dobet/models"

func CustomBet(owner_id string, match_id []string, odd float64, amount float64, market []models.Market, options []models.BetOption) models.Bet {
	bet := models.Bet{
		Bet_owner:     owner_id,
		Match_id:      match_id,
		Amount:        amount,
		Markets:       market,
		Options:       options,
		Odd:           odd,
		Potencial_win: odd * amount,
		IsProcessed:   false,
		IsFinished:    false,
	}
	return bet
}
