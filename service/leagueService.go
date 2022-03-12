package service

import (
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var LocalLeagues = make(map[string]models.League)
var LeagueService = &leagueService{}

type leagueService struct {
	repo        repository.LeagueRepository
	footballapi api.FootBallApi
}

func SetupLeagueService(leaguerepositorry repository.LeagueRepository, footballapi api.FootBallApi) {
	LeagueService.repo = leaguerepositorry
	LeagueService.footballapi = footballapi
	lunchUpdateLeaguesLoop()

}

func (service *leagueService) AddLeague(league models.League) error {
	return service.repo.AddLeague(league)
}

func (service *leagueService) DeleteLeague(league_id string) error {
	return service.repo.DeleteLeague(league_id)
}

func (service *leagueService) Leagues(startIndex, perpage int64) ([]models.League, error) {
	return service.repo.Leagues(startIndex, perpage)
}

func (service *leagueService)GetLeaguesByCountry(country string, startIndex, perpage int64) ([]models.League, error){
	return service.repo.GetLeaguesByCountry(country,startIndex,perpage)
}

func lunchUpdateLeaguesLoop() {
	tiker := time.NewTicker(time.Hour * 24 * 30)
	wg := &sync.WaitGroup{}

	for range tiker.C {
		if len(LocalLeagues) == 0 {
			localLeaguesdto, err := LeagueService.footballapi.GetLeagues()
			if err != nil {
				return
			}
			leagues := ConvertLeagueDtoToLeagueModelObjects(localLeaguesdto)
			requireGourotines := len(leagues)
			wg.Add(requireGourotines)
			for _, league := range leagues {
				go func(league models.League, wg *sync.WaitGroup) {
					defer wg.Done()
					LeagueService.AddLeague(league)
				}(league, wg)
			}
			for k := range LocalLeagues {
				delete(LocalLeagues, k)
			}
			for _, localleague := range leagues {
				LocalLeagues[localleague.League_id] = localleague
			}

		}
	}
	wg.Wait()
}
