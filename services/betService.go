package services

import (
	"gitthub.com/dionisiopro/dobet/helpers"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/observer"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var BetService betService
var BetProviders map[string]observer.BetProvider

type betService struct {
	repo repositories.BetRepository
}

func (service *betService) SetupBetService(betRepo repositories.BetRepository) *betService {
	BetService.repo = betRepo
	return &BetService
}

func (service *betService) CreateBet(owner_id string, match_id []string, odd float64, amount float64, market []models.Market, options []models.BetOption) error {
	bet := helpers.CustomBet(owner_id, match_id, odd, amount, market, options)
	bet_id, err := service.repo.CreateBet(bet)

	if err != nil {
		return err
	}

	consumer := helpers.CreateBetConsumer(bet_id)
	for _, provider := range BetProviders {
		for _, value := range match_id {
			if provider.Match_id == value {
				provider.AddConsumer(consumer)
			}
		}
	}
	return nil
}

func (service *betService) BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error) {
	return service.repo.BetByUser(user_id, startIndex, perpage)
}
func (service *betService) BetById(bet_id string) (models.Bet, error) {
	return service.repo.BetById(bet_id)
}

func (service *betService) BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error) {
	return service.repo.BetByUser(match_id, startIndex, perpage)
}

func (service *betService) Bets(startIndex, perpage int64) ([]models.Bet, error) {
	return service.repo.Bets(startIndex, perpage)
}

func (service *betService) RunningBets(startIndex, perpage int64) ([]models.Bet, error) {
	return service.repo.RunningBets(startIndex, perpage)
}

func (service *betService) TotalBets() (int, error) {
	return service.repo.TotalBets()
}
func (service *betService) TotalRunningBets() (int, error) {
	return service.repo.TotalRunningBets()
}

func (service *betService) TotalRunningBetsMoney() float64 {
	return service.repo.TotalRunningBetsMoney()
}

func  (service *betService)ProcessBet(bet_id string, match_result models.Match_Result) {
	bet, err := service.repo.BetById(bet_id)
	if err == nil && !bet.IsProcessed {
		for _, option := range bet.Options {
			if match_result.Match_id == option.Match_id && match_result.IsMatchFinished {
				if helpers.CheckWin(bet, match_result) {
					bet.IsLose = false
				}
				bet.RemainMatches--
				if bet.RemainMatches == 0 && !bet.IsLose {
					bet.IsFinished = true
					service.ProcessWin(bet.Potencial_win, bet.Bet_owner)
					bet.IsProcessed = true
				}
			}

		}
	}

}

func (service *betService) ProcessWin(amount float64, user_id string) {
	service.repo.ProcessWin(amount, user_id)
}


