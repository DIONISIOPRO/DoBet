package service

import (
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
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
	repository  repository.TeamRepository
	footballapi api.FootBallApi
}

func NewTeamService(teamRepository repository.TeamRepository, footballapi api.FootBallApi)TeamService {
return &teamService{
	repository: teamRepository,
	footballapi: footballapi,
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
	tiker := time.NewTicker(time.Hour * 24)
	wg := &sync.WaitGroup{}
	for range tiker.C {
		for _, league := range LocalLeagues {
			teamdto, err := service.footballapi.GetTeamsByLeagueId(league.League_id)
			if err != nil {
				return
			}
			teams := ConvertTeamDtoToTeamModelsObjects(teamdto)
			requiredGoroutines := len(teams)
			wg.Add(requiredGoroutines)
			for _, team := range teams {
				go func(team models.Team, wg *sync.WaitGroup) {
					defer wg.Done()
					service.Upsert(team)
				}(team, wg)
			}

		}
	}

}
