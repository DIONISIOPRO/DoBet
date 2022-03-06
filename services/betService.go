package services

import (
	"gitthub.com/dionisiopro/dobet/helpers"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var BetProviders map[string]BetProvider

type betService struct{
	betRepository repositories.BetRepository
}
func NewBetService(betRepository repositories.BetRepository) BetService{
	return &betService{
		betRepository: betRepository,
	}
}

func (betservice *betService) CreateBet(owner_id string, match_id []string, odd float64, amount float64, market []models.Market, options []models.BetOption) error {
	bet := helpers.CustomBet(owner_id, match_id, odd, amount, market, options)
	bet_id, err := betservice.betRepository.CreateBet(bet)

	if err != nil {
		return err
	}

	consumer := betservice.CreateBetConsumer(bet_id)
	for _, provider := range BetProviders {
		for _, value := range match_id {
			if provider.Match_id == value {
				provider.AddConsumer(consumer)
			}
		}
	}
	return nil
}

func(betservice *betService)  BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error) {
	return betservice.betRepository.BetByUser(user_id,startIndex, perpage)
}
func(betservice *betService) BetById(bet_id string) (models.Bet, error){
	return betservice.betRepository.BetById(bet_id)
}


func (betservice *betService) BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error){
	return betservice.betRepository.BetByUser(match_id,startIndex, perpage)
}

func (betservice *betService) Bets(startIndex, perpage int64) ([]models.Bet, error) {
	return betservice.betRepository.Bets(startIndex, perpage)
}

func(betservice *betService)  RunningBets(startIndex, perpage int64) ([]models.Bet, error) {
	return betservice.betRepository.RunningBets(startIndex, perpage)
}

func (betservice *betService) TotalRunningBetsMoney() float64 {
	return betservice.betRepository.TotalRunningBetsMoney()
}

func  (betservice *betService) CreateBetConsumer(bet_id string) BetConsumer {
	consumer := BetConsumer{
		BetId: bet_id,
		service: betservice,
	}

	return consumer
}

func  (betservice *betService) CreateBetProvider(match_id string) BetProvider {
	provider := BetProvider{
		Match_id:  match_id,
		Consumers: map[string]IBetConsumer{},
	}
	return provider
}

func (betservice *betService) ProcessBet(bet_id string, match_result models.Match_Result) {
	bet, err := betservice.betRepository.BetById(bet_id)
	if err == nil && !bet.IsProcessed {
		for _, option := range bet.Options {
			if match_result.Match_id == option.Match_id && match_result.IsMatchFinished {
				if CheckWin(bet, match_result){
					bet.IsLose = false
				}
				bet.RemainMatches--
				if bet.RemainMatches == 0 && !bet.IsLose {
					bet.IsFinished = true
					betservice.ProcessWin(bet.Potencial_win, bet.Bet_owner)
					bet.IsProcessed = true
				}
			}
 
		}
	}

}

func (betservice *betService) ProcessWin(amount float64, user_id string) {
	betservice.betRepository.ProcessWin(amount, user_id)
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
