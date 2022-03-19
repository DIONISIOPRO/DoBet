package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var BetProviders = map[string]models.BetProvider{}

type BetService interface {
	CreateBet(bet models.Bet) (string, error)
	BetByUser(user_id string, page, perpage int64) ([]models.Bet, error)
	BetById(bet_id string) (models.Bet, error)
	BetByMatch(match_id string, page, perpage int64) ([]models.Bet, error)
	Bets(page, perpage int64) ([]models.Bet, error)
	RunningBets(page, perpage int64) ([]models.Bet, error)
	TotalBets() (int, error)
	TotalRunningBets() (int, error)
	TotalRunningBetsMoney() float64
	ProcessBet(bet_id string, match_result models.Match_Result) error
}
type betService struct {
	repository repository.BetRepository
}

func NewBetService(betrepository repository.BetRepository) BetService {
	return &betService{
		repository: betrepository,
	}
}

func (service *betService) CreateBet(bet models.Bet) (string, error) {
	validate := validator.New()
	err := validate.Struct(bet)
	if err != nil {
		return "",err
	}
	bet_id, err := service.repository.CreateBet(bet)
	if err != nil {
		return"", err
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
	return bet_id, nil
}

func (service *betService) BetByUser(user_id string, page, perpage int64) ([]models.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if user_id == "" {
		return []models.Bet{}, errors.New("invalid user id")
	}
	return service.repository.BetByUser(user_id, startIndex, perpage)
}
func (service *betService) BetById(bet_id string) (models.Bet, error) {
	if bet_id == "" {
		return models.Bet{}, errors.New("invalid bet id")
	}
	return service.repository.BetById(bet_id)
}

func (service *betService) BetByMatch(match_id string, page, perpage int64) ([]models.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if match_id == "" {
		return []models.Bet{}, errors.New("invalid match id")
	}
	return service.repository.BetByUser(match_id, startIndex, perpage)
}

func (service *betService) Bets(page, perpage int64) ([]models.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Bets(startIndex, perpage)
}

func (service *betService) RunningBets(page, perpage int64) ([]models.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
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
	if bet_id == "" {
		return errors.New("invalid bet id")
	}
	validate := validator.New()
	err := validate.Struct(match_result)
	if err != nil {
		return err
	}
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
