package service

import (
	"strconv"
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
	LunchNewGoroutineToUpdateMatch(match models.Match, wg *sync.WaitGroup)
	LunchUpdateMatchesLoop()
	LunchProcessOldMatchesLoop()
}
type matchService struct {
	betService  BetService
	oddService  OddService
	repository  repository.MatchRepository
	footballApi api.FootBallApi
}

func NewMatchService(matchRepository repository.MatchRepository, betService BetService,
	footballApi api.FootBallApi, oddService OddService) MatchService {
	return &matchService{
		betService:  betService,
		repository:  matchRepository,
		footballApi: footballApi,
		oddService:  oddService,
	}
}

func (service *matchService) UpDateMatches(matchid string) error {
	matchdto, err := service.footballApi.GetNext20MatchesByLeagueId(matchid)
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
	matchdto, err := service.footballApi.GetLast5MatchesByLeagueId(match_id)
	if err != nil {
		return err
	}
	matches := ConvertMatchDtoToMatchModelsWithoutOddsObjects(matchdto)
	var wg = &sync.WaitGroup{}
	var totalrequeredGoroutines = len(matches)
	wg.Add(totalrequeredGoroutines)
	for _, match := range matches {
		go service.LunchNewGoroutineToDeleteMatch(match, wg)
	}
	wg.Wait()
	delete(BetProviders, match_id)
	return nil
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

func (service *matchService) LunchNewGoroutineToUpdateMatch(match models.Match, wg *sync.WaitGroup) {
	odd, err := service.oddService.GetOddByMatchId(match.Match_id)
	if err != nil {
		panic(err)
	}
	match.Odds = odd
	provider := CreateBetProvider(match.Match_id)
	BetProviders[match.Match_id] = provider
	service.repository.UpDateMatch(match.Match_id, match)
	if err != nil {
		wg.Done()
		panic(err)
	}

	wg.Done()
}

func (service *matchService) LunchNewGoroutineToDeleteMatch(match models.Match, wg *sync.WaitGroup) {
	if match.Result.IsMatchFinished {
		for _, provider := range BetProviders {
			if provider.Match_id == match.Match_id {
				provider.NotifyAll(match.Result, service.betService.ProcessBet)
			}
		}
	}
	id, _ := strconv.Atoi(match.Match_id)
	service.repository.DeleteOldMatchinCache(id)

	wg.Done()
}

func (service *matchService) LunchUpdateMatchesLoop() {
	tiker := time.NewTicker(time.Minute * 10)
	for range tiker.C {
		wg := &sync.WaitGroup{}
		wg.Add(len(LocalLeagues))
		requestMade := 0
		for _, id := range RequiredLeagueId {
			if requestMade%4 == 0 {
				time.Sleep(time.Second * 1)
			}
			go func(leagueId int64, wg *sync.WaitGroup) {
				id := strconv.Itoa(int(leagueId))
				service.UpDateMatches(id)
				defer wg.Done()
			}(id, wg)
			requestMade++
		}

	}
}

func (service *matchService) LunchProcessOldMatchesLoop() {
	tiker := time.NewTicker(time.Minute * 2)
	for range tiker.C {
		wg := &sync.WaitGroup{}
		requestMade := 0
		wg.Add(len(RequiredLeagueId))
		for _, id := range RequiredLeagueId {
			if requestMade%6 == 0 {
				time.Sleep(time.Second * 1)
			}
			go func(leagueId int64, wg *sync.WaitGroup) {
				id := strconv.Itoa(int(leagueId))
				go service.DeleteOldMatch(id)
				defer wg.Done()
			}(id, wg)
			requestMade++
		}

	}
}
