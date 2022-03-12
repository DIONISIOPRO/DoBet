package service

import (
	"log"
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var MatchService = &matchService{}

type matchService struct {
	repository  repository.MatchRepository
	footballApi api.FootBallApi
}

func SetupMatchService(matchRepository repository.MatchRepository, footballApi api.FootBallApi) {
	MatchService.repository = matchRepository
	MatchService.footballApi = footballApi
	MatchService.MatchWatch()
	lunchUpdateMatchesLoop()
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
		go lunchNewGoroutineToUpdateMatch(match, wg)
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

func (service *matchService) Matches(startIndex, perpage int64) ([]models.Match, error) {
	return service.repository.Matches(startIndex, perpage)
}
func (service *matchService) MatchesByLeagueId(leagueid string, startIndex, perpage int64) ([]models.Match, error) {
	return service.repository.MatchesByLeagueId(leagueid,startIndex, perpage)
}

func (service *matchService) MatchesByLeagueIdDay(leagueid string, day, startIndex, perpage int64) ([]models.Match, error){
	return service.repository.MatchesByLeagueIdDay(leagueid, day,startIndex, perpage)
	
}

func (service *matchService) MatchWatch() {
	go service.repository.MatchWatch(onMatchStreamChange)
}

func onMatchStreamChange(match models.Match) {
	result := match.Result
	for _, betprovider := range BetProviders {
		if match.Match_id == betprovider.Match_id {
			go betprovider.NotifyAll(result, BetService.ProcessBet)
		}
	}
}

func lunchNewGoroutineToUpdateMatch(match models.Match, wg *sync.WaitGroup) {
	odd := OddService.GetOddByMatchId(match.Match_id)
	match.Odds = odd
	err := MatchService.repository.UpDateMatch(match.Match_id, match)
	if err != nil {
		wg.Done()
		panic(err)
	}
	provider := CreateBetProvider(match.Match_id)
	BetProviders[match.Match_id] = provider
	wg.Done()
}

func lunchUpdateMatchesLoop() {
	tiker := time.NewTicker(time.Minute * 2)
	for tker := range tiker.C {
		wg := &sync.WaitGroup{}
		log := log.Default()
		log.Print(tker)
		wg.Add(len(LocalLeagues))
		for _, league := range LocalLeagues {
			go func(leagueId string, wg *sync.WaitGroup){
				go MatchService.UpDateMatches(leagueId)		
				defer wg.Done()	

			}(league.League_id,wg)
		}
		
	}
}

