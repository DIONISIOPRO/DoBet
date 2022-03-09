package service

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/observer"
	"gitthub.com/dionisiopro/dobet/repository"
)

var BetService = &betService{}
var BetProviders = map[string]observer.BetProvider{}

type betService struct {
	repository repository.BetRepository
}

func (service *betService) SetupBetService(betrepository repository.BetRepository) {
	BetService.repository = betrepository
}

func (service *betService) CreateBet(bet models.Bet) error {
	bet_id, err := service.repository.CreateBet(bet)
	if err != nil {
		return err
	}
	consumer := CreateBetConsumer(bet_id)
	betsId := bet.BetGroup
	for _, provider := range BetProviders {
		for _, value := range betsId {
			if provider.Match_id == value.Match_id {
				provider.AddConsumer(consumer)
			}
		}
	}
	return nil
}

func (service *betService) BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error) {
	return service.repository.BetByUser(user_id, startIndex, perpage)
}
func (service *betService) BetById(bet_id string) (models.Bet, error) {
	return service.repository.BetById(bet_id)
}

func (service *betService) BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error) {
	return service.repository.BetByUser(match_id, startIndex, perpage)
}

func (service *betService) Bets(startIndex, perpage int64) ([]models.Bet, error) {
	return service.repository.Bets(startIndex, perpage)
}

func (service *betService) RunningBets(startIndex, perpage int64) ([]models.Bet, error) {
	return service.repository.RunningBets(startIndex, perpage)
}

func (service *betService) TotalBets() (int, error) {
	return service.repository.TotalBets()
}
func (service *betService) TotalRunningBets() (int, error) {
	return service.repository.TotalRunningBets()
}

func (service *betService) TotalRunningBetsMoney() float64 {
	return service.repository.TotalRunningBetsMoney()
}

func (service *betService) ProcessBet(bet_id string, match_result models.Match_Result) error {
	bet, err := service.repository.BetById(bet_id)
	if err != nil {
		return err
	}

	if !bet.IsFinished {
		matchId := match_result.Match_id
		for _, bet := range bet.BetGroup {
			if !bet.IsProcessed && bet.Match_id == matchId {
				switch bet.Market.(type) {
				case models.AllScoreMarket:
					localbet := bet.Market.(models.AllScoreMarket)
					if localbet.Option != match_result.All_Scores {
						bet.IsLose = true
					}
				case models.WinnerMarket:
					localbet := bet.Market.(models.WinnerMarket)
					if localbet.Option != match_result.Winner {
						bet.IsLose = true
					}
				}
			}
			bet.IsProcessed = true

		}

		for _, localbet := range bet.BetGroup {
			if !localbet.IsProcessed {
				return nil
			}
			if localbet.IsLose {
				return nil
			}
		}
		err = service.repository.UpdateBet(bet.Bet_id, bet)
		if err != nil {
			return err
		}
		service.repository.ProcessWin(bet.TotalAmount, bet.Bet_owner)
		bet.IsFinished = true
		return nil
	}
	return nil

}
