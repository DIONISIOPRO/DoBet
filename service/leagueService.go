package service

import (
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var LocalLeagues = make(map[string]models.League)

type LeagueService interface {
	AddLeague(league models.League) error
	DeleteLeague(league_id string) error
	Leagues(page, perpage int64) ([]models.League, error)
	GetLeaguesByCountry(country string, page, perpage int64) ([]models.League, error)
	LunchUpdateLeaguesLoop() 
}
type leagueService struct {
	repo        repository.LeagueRepository
	footballapi api.FootBallApi
}

func NewLeagueService(leaguerepositorry repository.LeagueRepository, footballapi api.FootBallApi) LeagueService {
	return &leagueService{
		repo:        leaguerepositorry,
		footballapi: footballapi,
	}
}

func (service *leagueService) AddLeague(league models.League) error {
	return service.repo.AddLeague(league)
}

func (service *leagueService) DeleteLeague(league_id string) error {
	return service.repo.DeleteLeague(league_id)
}

func (service *leagueService) Leagues(page, perpage int64) ([]models.League, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repo.Leagues(startIndex, perpage)
}

func (service *leagueService) GetLeaguesByCountry(country string, page, perpage int64) ([]models.League, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repo.GetLeaguesByCountry(country, startIndex, perpage)
}

func(service *leagueService) LunchUpdateLeaguesLoop() {
	tiker := time.NewTicker(time.Hour * 24 * 30)
	wg := &sync.WaitGroup{}

	for range tiker.C {
		if len(LocalLeagues) == 0 {
			localLeaguesdto, err := service.footballapi.GetLeagues()
			if err != nil {
				return
			}
			leagues := ConvertLeagueDtoToLeagueModelObjects(localLeaguesdto)
			requireGourotines := len(leagues)
			wg.Add(requireGourotines)
			for _, league := range leagues {
				go func(league models.League, wg *sync.WaitGroup) {
					defer wg.Done()
					service.AddLeague(league)
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
