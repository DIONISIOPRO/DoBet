package service

import (
	"strconv"
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/data"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

type TeamService interface {
	Upsert(team models.Team) error
	DeleteTeam(team_id string) error
	Teams(page, perpage int64) ([]models.Team, error)
	LunchUpdateTeamssLoop()
	TeamsByCountry(country string, page, perpage int64) ([]models.Team, error)
}

type teamService struct {
	repository   repository.TeamRepository
	footballdata data.FootballData
}

func NewTeamService(teamRepository repository.TeamRepository, footballdata data.FootballData) TeamService {
	return &teamService{
		repository:   teamRepository,
		footballdata: footballdata,
	}
}

func (service *teamService) Upsert(team models.Team) error {
	return service.repository.Upsert(team)
}

func (service *teamService) DeleteTeam(team_id string) error {
	return service.repository.DeleteTeam(team_id)
}

func (service *teamService) Teams(page, perpage int64) ([]models.Team, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.Teams(startIndex, perpage)
}

func (service *teamService) TeamsByCountry(country string, page, perpage int64) ([]models.Team, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 9
	}
	startIndex := (page - 1) * perpage
	return service.repository.TeamsByCountry(country, startIndex, perpage)
}

func (service *teamService) LunchUpdateTeamssLoop() {
	tiker := time.NewTicker(time.Hour * 24 * 7)
	wg := &sync.WaitGroup{}
	for _, leagueId := range RequiredLeagueId {
		time.Sleep(time.Minute * 4)
		id := strconv.Itoa(int(leagueId))
		teams, err := service.footballdata.GetTeamsByLeagueId(id)
		if err != nil {
			return
		}
		requiredGoroutines := len(teams)
		wg.Add(requiredGoroutines)
		for _, team := range teams {
			go func(team models.Team, wg *sync.WaitGroup) {
				defer wg.Done()
				service.Upsert(team)
			}(team, wg)
		}
	}
	for range tiker.C {
		service.LunchUpdateTeamssLoop()
	}
	wg.Wait()

}
