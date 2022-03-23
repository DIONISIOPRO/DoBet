package service

import (
	"gitthub.com/dionisiopro/dobet/data"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var LocalLeagues = make(map[string]models.League)
var RequiredLeagueId = []int64{128, 129, 71, 75, 72, 88, 89, 144, 235, 106, 61, 79, 78, 62, 140, 141, 80, 145, 107, 95, 94, 135, 136}

type LeagueService interface {
	AddManyLeague(league []models.League) error
	DeleteLeague(league_id string) error
	Leagues(page, perpage int64) ([]models.League, error)
	GetLeaguesByCountry(country string, page, perpage int64) ([]models.League, error)
	LunchUpdateLeaguesLoop()
}
type leagueService struct {
	repo        repository.LeagueRepository
	footballdata data.FootballData
}

func NewLeagueService(leaguerepositorry repository.LeagueRepository, footballdata data.FootballData) LeagueService {
	return &leagueService{
		repo:        leaguerepositorry,
		footballdata: footballdata,
	}
}

func (service *leagueService) AddManyLeague(league []models.League) error {
	return service.repo.AddManyLeague(league)
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

func (service *leagueService) LunchUpdateLeaguesLoop() {
	// tiker := time.NewTicker(time.Second * 5)
	for _, id := range RequiredLeagueId {
		league, err := service.footballdata.GetLeague(id)
		if err != nil {
			return
		}
		leagues := []models.League{}
		leagues = append(leagues, league)
		go service.repo.AddManyLeague(leagues)
	}

	// for range tiker.C {
	// 	log.Println("TICKER DISPACHER : starting to fech api to get leagues dto")
	// 	service.LunchUpdateLeaguesLoop()
	// }
}
