package service

import (
	"errors"

	"github.com/dionisiopro/dobet-bet/domain/bet"
	"github.com/dionisiopro/dobet-bet/domain/event"
	"github.com/dionisiopro/dobet-bet/domain/result"
)

type BetRepository interface {
	ConfirmBet(bet_id string) error
	ActiveBet(bet_id string) error
	CancelBet(bet_id string) error
	CreateBet(bet bet.BetBase) (bet_id string, err error)
	UpdateBet(bet_id string, bet bet.BetBase) error
	BetByUser(user_id string, startIndex, perpage int64) ([]bet.BetBase, error)
	BetByMatch(match_id string, startIndex, perpage int64) ([]bet.BetBase, error)
	AllRunningBetsByMatch(match_id string) ([]bet.BetBase, error)
	BetById(bet_id string) (bet.BetBase, error)
	TotalBets() (int, error)
	TotalRunningBets() (int, error)
	Bets(startIndex, perpage int64) ([]bet.BetBase, error)
	RunningBets(startIndex, perpage int64) ([]bet.BetBase, error)
	TotalRunningBetsMoney() float64
}
type Event interface {
	ToByteArray() ([]byte, error)
}
type EventPublisher interface {
	Publish(topic string, event Event) error
}

type BetService struct {
	eventPublisher EventPublisher
	repository     BetRepository
}

func NewBetService(betrepository BetRepository, eventPublisher EventPublisher) *BetService {
	return &BetService{
		repository:     betrepository,
		eventPublisher: eventPublisher,
	}
}

func (service *BetService) CreateBet(bet *bet.BetBase) (string, error) {
	if !bet.IsValid() {
		return "", errors.New("bet invalid")
	}
	bet.Status = "created"
	for _, _bet := range bet.BetGroup {
		_bet.Result = result.MatchResult{}
	}
	id, err := service.repository.CreateBet(*bet)
	if err != nil {
		return "", err
	}
	var ids []string
	for _, _bet := range bet.BetGroup {
		ids = append(ids, _bet.Match_id)
	}

	betCreatedEvent := event.BetCreatedEvent{
		User_id:   bet.Bet_owner,
		Bet_id:    bet.Bet_id,
		Match_idS: ids,
	}
	service.eventPublisher.Publish(event.BetCreated, betCreatedEvent)
	return id, nil
}

func (service *BetService) BetByUser(user_id string, page, perpage int64) ([]bet.BetBase, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if user_id == "" {
		return []bet.BetBase{}, errors.New("invalid user id")
	}
	return service.repository.BetByUser(user_id, startIndex, perpage)
}

func (service *BetService) BetById(bet_id string) (bet.BetBase, error) {
	if bet_id == "" {
		return bet.BetBase{}, errors.New("invalid bet id")
	}
	return service.repository.BetById(bet_id)
}

func (service *BetService) BetByMatch(match_id string, page, perpage int64) ([]bet.BetBase, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	if match_id == "" {
		return []bet.BetBase{}, errors.New("invalid match id")
	}
	return service.repository.BetByUser(match_id, startIndex, perpage)
}

func (service *BetService) Bets(page, perpage int64) ([]bet.BetBase, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Bets(startIndex, perpage)
}

func (service *BetService) RunningBets(page, perpage int64) ([]bet.BetBase, error) {
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

func (service *BetService) ConfirmBet(bet_id string) error {
	return service.repository.ConfirmBet(bet_id)

}
func (service *BetService) ActiveBet(bet_id string) error {
	return service.repository.ActiveBet(bet_id)
}
func (service *BetService) CancelBet(bet_id string) error {
	return service.CancelBet(bet_id)
}

func (service *BetService) ProcessMatchResultInBet(result result.MatchResult) error {
	bets, err := service.repository.AllRunningBetsByMatch(result.Match_id)
	if err != nil {
		return err
	}
	betLenc := len(bets)
	betChann := make(chan bet.BetBase, betLenc+1)
	for _, bet := range bets {
		go updateBet(&bet, result, betChann)
	}
	for i := 0; i < betLenc; i++ {
		_bet := <-betChann
		service.repository.UpdateBet(_bet.Bet_id, _bet)
		service.finishBet(&_bet)
	}
	return nil
}

func (service *BetService) finishBet(bet *bet.BetBase) {
	if !bet.GetIsFinished() {
		return
	}
	if bet.IsLose() {
		return
	}
	depositEvent := event.BetDepositEvent{
		User_id: bet.Bet_owner,
		Amount:  bet.GetPotenctialWin(),
	}
	service.eventPublisher.Publish(event.BetDeposit, depositEvent)
	bet.IsFinished = true
	service.repository.UpdateBet(bet.Bet_id, *bet)
}

func updateBet(bet *bet.BetBase, result result.MatchResult, betChann chan bet.BetBase) {
	for index, _bet := range bet.BetGroup {
		if _bet.Match_id != result.Match_id {
			continue
		}
		_bet.SetResult(result)
		if _bet.GetIsLose() {
			bet.BetGroup[index].IsLose = true
		}
		bet.BetGroup[index].IsProcessed = true
	}
	betChann <- *bet
}
