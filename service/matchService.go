package service

import (
	"log"
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

type MatchService interface {
	UpDateMatches(matchid string) error
	DeleteOldMatch(match_id string) error
	MatchesByLeagueIdDay(leagueid string, day, page, perpage int64) ([]models.Match, error)
	MatchWatch()
	LunchNewGoroutineToUpdateMatch(match models.Match, wg *sync.WaitGroup)
	LunchUpdateMatchesLoop()
}
type matchService struct {
	betService  BetService
	oddService OddService
	repository  repository.MatchRepository
	footballApi api.FootBallApi
}

func NewMatchService(matchRepository repository.MatchRepository, betService BetService,
	footballApi api.FootBallApi, 	oddService OddService) MatchService {
	return &matchService{
		betService:  betService,
		repository:  matchRepository,
		footballApi: footballApi,
		oddService: oddService,
	}
}

func (service *matchService) UpDateMatches(matchid string) error {
	matchdto, err := service.footballApi.GetMatchesByLeagueId(matchid)
	if err != nil {
		return err
	}
	matches := ConvertMatchDtoToMatchModelsWithoutOddsObjects(matchdto)
	var wg = &sync.WaitGroup{}
	var totalrequeredGoroutines = len(matches)
	wg.Add(totalrequeredGoroutines)
	for _, match := range matches {
		go service.LunchNewGoroutineToUpdateMatch(match, wg)
	}
	wg.Wait()
	return nil
}

func (service *matchService) DeleteOldMatch(match_id string) error {
	err := service.repository.DeleteOldMatch()
	if err != nil {
		return err
	}
	ConsumersLength := len(BetProviders[match_id].Consumers)
	if ConsumersLength != 0 {
		return nil
	}
	delete(BetProviders, match_id)
	return nil
}

func (service *matchService) Matches(page, perpage int64) ([]models.Match, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Matches(startIndex, perpage)
}
func (service *matchService) MatchesByLeagueId(leagueid string, page, perpage int64) ([]models.Match, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.MatchesByLeagueId(leagueid, startIndex, perpage)
}

func (service *matchService) MatchesByLeagueIdDay(leagueid string, day, page, perpage int64) ([]models.Match, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.MatchesByLeagueIdDay(leagueid, day, startIndex, perpage)

}

func (service *matchService) MatchWatch() {
	go service.repository.MatchWatch(service.OnMatchStreamChange)
}

func (service *matchService) OnMatchStreamChange(match models.Match) {
	result := match.Result
	for _, betprovider := range BetProviders {
		if match.Match_id == betprovider.Match_id {
			go betprovider.NotifyAll(result, service.betService.ProcessBet)
		}
	}
}

func (service *matchService) LunchNewGoroutineToUpdateMatch(match models.Match, wg *sync.WaitGroup) {
	odd, err := service.oddService.GetOddByMatchId(match.Match_id)
	if  err != nil {
		panic(err)
	}
	match.Odds = odd
	err = service.repository.UpDateMatch(match.Match_id, match)
	if err != nil {
		wg.Done()
		panic(err)
	}
	provider := CreateBetProvider(match.Match_id)
	BetProviders[match.Match_id] = provider
	wg.Done()
}

func (service *matchService) LunchUpdateMatchesLoop() {
	tiker := time.NewTicker(time.Minute * 2)
	for tker := range tiker.C {
		wg := &sync.WaitGroup{}
		log := log.Default()
		log.Print(tker)
		wg.Add(len(LocalLeagues))
		for _, league := range LocalLeagues {
			go func(leagueId string, wg *sync.WaitGroup) {
				go service.UpDateMatches(leagueId)
				defer wg.Done()

			}(league.League_id, wg)
		}

	}
}
