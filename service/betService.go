package service

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var BetProviders = map[string]models.BetProvider{}

type BetService interface {
	CreateBet(bet *models.Bet) (string, error)
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

func (service *betService) CreateBet(bet *models.Bet) (string, error) {
	validate := validator.New()
	err := validate.Struct(bet)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var odd float64
	for _, localbet := range bet.BetGroup {
		odd = odd + localbet.Odd
		localbet.IsLose = false
		localbet.IsProcessed = false
		switch localbet.Market {
		case "WINNER":
			if localbet.Option.Will_Team_Away_wins {
				if localbet.Option.Will_Team_Home_wins || localbet.Option.Will_Draw {
					return "", errors.New("cannot select more than one option to make bet in winner")
				}
			} else if localbet.Option.Will_Team_Home_wins {
				if localbet.Option.Will_Team_Away_wins || localbet.Option.Will_Draw {
					return "", errors.New("cannot select more than one option to make bet in winner")
				}
			} else if localbet.Option.Will_Draw {
				if localbet.Option.Will_Team_Away_wins || localbet.Option.Will_Draw {
					return "", errors.New("cannot select more than one option to make bet in winner")
				}
			}

		}
	}
	bet.GlobalOdd = float64(odd)
	bet.IsFinished = false
	if bet.Potencial_win != bet.TotalAmount*bet.GlobalOdd {
		return "", errors.New("the potencial win money dont match whith odd value and amount")
	}
	bet_id, err := service.repository.CreateBet(*bet)
	if err != nil {
		log.Println(err)
		return "", err
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
		log.Println(err)
		return err
	}
	bet, err := service.repository.BetById(bet_id)
	if err != nil {
		return err
	}

	if bet.IsFinished {
		return nil
	}
	matchId := match_result.Match_id
	for _, bet := range bet.BetGroup {
		if bet.IsProcessed || bet.Match_id != matchId {
			continue
		}
		switch bet.Market {
		case "ALLSCORE":
			if bet.Option.Will_All_Scores != match_result.All_Scores {
				bet.IsLose = true
			}
		case "WINNER":
			homewins := match_result.Team_Home_Goals > match_result.Team_Away_Goals
			awaywins := match_result.Team_Home_Goals < match_result.Team_Away_Goals
			wasdraw := match_result.Team_Away_Goals == match_result.Team_Home_Goals
			if bet.Option.Will_Draw && !wasdraw {
				bet.IsLose = true
			}
			if bet.Option.Will_Team_Away_wins && !awaywins {
				bet.IsLose = true
			}
			if bet.Option.Will_Team_Home_wins && homewins {
				bet.IsLose = true
			}
		}

		bet.IsProcessed = true

	}

	for _, localbet := range bet.BetGroup {
		// if all bet processed continue to is the uer wins
		if !localbet.IsProcessed {
			return nil
		}
		//if all bet wins proced to pocess win otherwise return here
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
