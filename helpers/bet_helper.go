package helpers

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/observer"



)

func CustomBet(
	owner_id string,
	match_id []string,
	odd float64,
	amount float64,
	market []models.Market,
	options []models.BetOption,
) models.Bet {

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


func  CreateBetConsumer(bet_id string) observer.BetConsumer {
	consumer := observer.BetConsumer{
		BetId: bet_id,
	}
	return consumer
}

func  CreateBetProvider(match_id string) observer.BetProvider {
	provider := observer.BetProvider{
		Match_id:  match_id,
		Consumers: map[string]observer.BetConsumer{},
	}
	return provider
}


func  CheckWin(bet models.Bet, match_result models.Match_Result) bool {
	IsWin := true
	for _, option := range bet.Options {
		if match_result.Match_id == option.Match_id && match_result.IsMatchFinished {

			if match_result.All_Scores && !option.Will_All_Scores {
				IsWin = false
			}

			if match_result.Is_Draw && !option.Will_Draw {
				IsWin = false
			}

			if match_result.Is_Team_Away_wins && !option.Will_Team_Away_wins {
				IsWin = false
			}

			if match_result.Is_Team_Home_wins && !option.Will_Team_Home_wins {
				IsWin = false
			}
		}
	}
	return IsWin
}




