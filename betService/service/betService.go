package service

import (
	"errors"

	"github.com/dionisiopro/dobet-bet/domain/bet"
	"github.com/dionisiopro/dobet-bet/domain/result"
)
type BetRepository interface {
	CreateBet(bet bet.BetBaseImpl) (bet_id string, err error)
	UpdateBet(bet_id string, bet bet.BetBaseImpl) error
	BetByUser(user_id string, startIndex, perpage int64) ([]bet.BetBaseImpl, error)
	BetByMatch(match_id string, startIndex, perpage int64) ([]bet.BetBaseImpl, error)
	AllRunningBetsByMatch(match_id string) ([]bet.BetBaseImpl, error)
	BetById(bet_id string) (bet.BetBaseImpl, error)
	TotalBets() (int, error)
	TotalRunningBets() (int, error)
	Bets(startIndex, perpage int64) ([]bet.BetBaseImpl, error)
	RunningBets(startIndex, perpage int64) ([]bet.BetBaseImpl, error)
	TotalRunningBetsMoney() float64
}


type BetService struct {
	repository BetRepository
}

func NewBetService(betrepository BetRepository) *BetService {
	return &BetService{
		repository: betrepository,
	}
}

func (service *BetService) CreateBet(bet *bet.BetBaseImpl) (string, error) {
	if !bet.IsValid(){
		return "", errors.New("bet invalid")
	}
	bet.Status = "created"
	for _, _bet := range bet.BetGroup{
		_bet.Result = nil
	}
	id, err := service.repository.CreateBet(*bet)
	if err != nil{
		return "", err
	}

	//TODO : publishing bet created event

	return id, nil
}

func (service *BetService) BetByUser(user_id string, page, perpage int64) ([]bet.BetBaseImpl, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if user_id == "" {
		return []bet.BetBaseImpl{}, errors.New("invalid user id")
	}
	return service.repository.BetByUser(user_id, startIndex, perpage)
}

func (service *BetService) BetById(bet_id string) (bet.BetBaseImpl, error) {
	if bet_id == "" {
		return bet.BetBaseImpl{}, errors.New("invalid bet id")
	}
	return service.repository.BetById(bet_id)
}

func (service *BetService) BetByMatch(match_id string, page, perpage int64) ([]bet.BetBaseImpl, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if match_id == "" {
		return []bet.BetBaseImpl{}, errors.New("invalid match id")
	}
	return service.repository.BetByUser(match_id, startIndex, perpage)
}

func (service *BetService) Bets(page, perpage int64) ([]bet.BetBaseImpl, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Bets(startIndex, perpage)
}

func (service *BetService) RunningBets(page, perpage int64) ([]bet.BetBaseImpl, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.RunningBets(startIndex, perpage)
}

func (service *BetService) TotalBets() (int, error) {
	return service.repository.TotalBets()
}

func (service *BetService) TotalRunningBets() (int, error) {
	return service.repository.TotalRunningBets()
}

func (service *BetService) TotalRunningBetsMoney() float64 {
	return service.repository.TotalRunningBetsMoney()
}

func (service *BetService) UpdateBetByMatchResult(result result.MatchResultImpl) error {
	bets, err := service.repository.AllRunningBetsByMatch(result.Match_id)
	if err != nil {
		return err
	}
	betLenc := len(bets)
	betChann :=  make(chan bet.BetBaseImpl, betLenc + 1)
	for _, bet := range bets{
		go	updateBet(&bet, result, betChann)
	}
	for i := 0; i < betLenc; i++{
		_bet :=  <- betChann
        service.repository.UpdateBet(_bet.Bet_id, _bet)
		service.finishBet(&_bet)
	}
	return nil
}

func (service *BetService) finishBet(bet *bet.BetBaseImpl){
	if !bet.GetIsFinished(){
		return
	}
	if bet.IsLose(){
		return
	}
	// TODO : publish user bet win event
	bet.IsFinished = true
	service.repository.UpdateBet(bet.Bet_id, *bet)
}

func updateBet(bet *bet.BetBaseImpl, result result.MatchResultImpl, betChann chan bet.BetBaseImpl){
	for index, _bet := range bet.BetGroup{
		if  _bet.Match_id != result.Match_id{
			continue
		}
		_bet.SetResult(result)
		if  _bet.GetIsLose() {
			bet.BetGroup[index].IsLose = true
		}
		bet.BetGroup[index].IsProcessed = true
	}
	betChann <- *bet
}
