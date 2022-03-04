package services

import (
	"gitthub.com/dionisiopro/dobet/helpers"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var BetProviders map[string]BetProvider

func CreateBet(owner_id string, match_id []string, odd float64, amount float64, market []models.Market, options []models.BetOption) error {
	bet := helpers.CustomBet(owner_id, match_id, odd, amount, market, options)
	bet_id, err := repositories.CreateBet(bet)

	if err != nil {
		return err
	}
	consumer := CreateBetConsumer(bet_id)
	for _, provider := range BetProviders {
		for _, value := range match_id {
			if provider.Match_id == value {
				provider.AddConsumer(consumer)
			}
		}
	}

	return nil

}

func BetByUser(user_id string) []models.Bet {
	return repositories.BetByUser(user_id)
}

func BetByMatch(match_id string) []models.Bet {
	return repositories.BetByUser(match_id)
}

func Bets() []models.Bet {
	return repositories.Bets()
}

func RunningBets() []models.Bet {
	return repositories.RunningBets()
}

func TotalRunningBetsMoney() float32 {
	return repositories.TotalRunningBetsMoney()
}

func CreateBetConsumer(bet_id string) BetConsumer {
	consumer := BetConsumer{
		BetId: bet_id,
	}

	return consumer
}

func CreateBetProvider(match_id string) BetProvider {
	provider := BetProvider{
		Match_id:  match_id,
		Consumers: map[string]IBetConsumer{},
	}

	return provider
}

func ProcessBet(bet_id string, match_result models.Match_Result) {
	bet := repositories.BetById(bet_id)
	if !bet.IsProcessed {
		for _, option := range bet.Options {
			if match_result.Match_id == option.Match_id && match_result.IsMatchFinished {
				if CheckWin(bet, match_result){
					bet.IsLose = false
				}
				bet.RemainMatches--
				if bet.RemainMatches == 0 && !bet.IsLose {
					bet.IsFinished = true
					ProcessWin(bet.Potencial_win, bet.Bet_owner)
					bet.IsProcessed = true
				}
			}

		}
		repositories.UpdateBet(bet_id, bet)
	}

}

func ProcessWin(amount float64, user_id string) {
	repositories.Win(amount, user_id)
}

func CheckWin(bet models.Bet, match_result models.Match_Result) bool {
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
