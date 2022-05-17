package service

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"gitthub.com/dionisiopro/dobet/domain"
	"gitthub.com/dionisiopro/dobet/repository"
)
type BetRepository interface {
	CreateBet(bet domain.Bet) (bet_id string, err error)
	UpdateBet(bet_id string, bet domain.Bet) error
	BetByUser(user_id string, startIndex, perpage int64) ([]domain.Bet, error)
	BetByMatch(match_id string, startIndex, perpage int64) ([]domain.Bet, error)
	AllRunningBetsByMatch(match_id string) ([]domain.Bet, error)
	BetById(bet_id string) (domain.Bet, error)
	TotalBets() (int, error)
	TotalRunningBets() (int, error)
	Bets(startIndex, perpage int64) ([]domain.Bet, error)
	RunningBets(startIndex, perpage int64) ([]domain.Bet, error)
	TotalRunningBetsMoney() float64
}
var BetProviders = map[string]domain.BetProvider{}


type betService struct {
	repository BetRepository
}

func NewBetService(betrepository BetRepository) BetService {
	return &betService{
		repository: betrepository,
	}
}

func (service *betService) CreateBet(bet *domain.Bet) (string, error) {
	if !bet.IsValid(){
		return "", errors.New("bet invalid")
	}
	bet.Status = "created"
	for _, _bet :=range bet.BetGroup{
		_bet.Result = nil
	}
	err, id := service.repository.CreateBet(bet)
	if err != nil{
		return "", err
	}

	//TODO : publishing bet created event

	return id, nil
}

func (service *betService) BetByUser(user_id string, page, perpage int64) ([]domain.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if user_id == "" {
		return []domain.Bet{}, errors.New("invalid user id")
	}
	return service.repository.BetByUser(user_id, startIndex, perpage)
}

func (service *betService) BetById(bet_id string) (domain.Bet, error) {
	if bet_id == "" {
		return domain.Bet{}, errors.New("invalid bet id")
	}
	return service.repository.BetById(bet_id)
}

func (service *betService) BetByMatch(match_id string, page, perpage int64) ([]domain.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if match_id == "" {
		return []domain.Bet{}, errors.New("invalid match id")
	}
	return service.repository.BetByUser(match_id, startIndex, perpage)
}

func (service *betService) Bets(page, perpage int64) ([]domain.Bet, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Bets(startIndex, perpage)
}

func (service *betService) RunningBets(page, perpage int64) ([]domain.Bet, error) {
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

func (service *betService) UpdateBetByMatchResult(result domain.MatchResultBase) error {
	bets, err := service.repository.AllRunningBetsByMatch(result.Match_id)
	if err != nil {
		return err
	}
	betLenc := len(bets)
	betChann :=  make(chan domain.BetBaseImpl, betLenc + 1)
	for _, bet := range bets{
		go	updateBet(bet, result, betChann)
	}
	for i := 0; i < betLenc; i++{
		_bet :=  <- betChann
        service.repository.UpdateBet(_bet.Bet_id, _bet)
		service.finishBet(_bet)
	}
}

func (service *betService) finishBet(bet *domain.BetBaseImpl){
	if !bet.IsFinished(){
		return
	}
	if bet.IsLose(){
		return
	}
	// TODO : publish user bet win event
	bet.IsFinished = true
	service.repository.UpdateBet(bet.Bet_id, bet)
}

func updateBet(bet *domain.BetBaseImpl, result domain.MatchResultBase, betChann chan domain.BetBaseImpl){
	for index, _bet := range bet.BetGroup{
		if  _bet.Match_id != result.match_id{
			continue
		}
		_bet.SetResult(result)
		if  _bet.IsLose() {
			bet[index].IsLose = true
		}
		bet[index].IsProcessed = true
	}
	betChann <- bet
}
