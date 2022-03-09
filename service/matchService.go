package service

import (
	"strconv"
	"sync"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var MatchService = &matchService{}

type matchService struct {
	repository        repository.MatchRepository
	footballApi api.FootBallApi
}

func SetupMatchService(matchRepository repository.MatchRepository, footballApi api.FootBallApi) {
	MatchService.repository = matchRepository
	MatchService.footballApi = footballApi
	MatchService.MatchWatch()
}

func (service *matchService) UpDateMatches() error {
	matchdto, err := service.footballApi.Matches()
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
	if  ConsumersLength != 0{
		return nil
	}
	delete(BetProviders, match_id)
	return nil
}

func (service *matchService) Matches(startIndex, perpage int64) ([]models.Match, error) {
	return service.repository.Matches(startIndex, perpage)
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

func lunchNewGoroutineToUpdateMatch(match models.Match, wg *sync.WaitGroup){
		matchId, err := strconv.Atoi(match.Match_id)
		if err != nil {
			wg.Done()
			panic(err)
		}
		odd := getOddsToAddInMatch(matchId, wg)
		match.Odds = odd
		err = MatchService.repository.UpDateMatch(match.Match_id, match)
		if err != nil {
			wg.Done()
			panic(err)
		}
		provider := CreateBetProvider(match.Match_id)
		BetProviders[match.Match_id] = provider
		wg.Done()
}

func getOddsToAddInMatch(matchId int, wg *sync.WaitGroup)models.Odds{
	odd_dto, err := MatchService.footballApi.GetOddsByMatchId(matchId)
	if err != nil {
		wg.Done()
		panic(err)
	}
	odd := ConvertOddDtoToOddModelObject(odd_dto)
	return odd
}
