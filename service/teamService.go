package service

import (
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var TeamService = &teamService{}

type teamService struct {
	repository  repository.TeamRepository
	footballapi api.FootBallApi
}

func SetupTeamService(teamRepository repository.TeamRepository, footballapi api.FootBallApi) {
	TeamService.repository = teamRepository
	TeamService.footballapi = footballapi
	lunchUpdateTeamssLoop()
}

func (service *teamService) Upsert(team models.Team) error {
	return service.repository.Upsert(team)
}

func (service *teamService) DeleteTeam(team_id string) error {
	return service.repository.DeleteTeam(team_id)
}

func (service *teamService) Teams(startIndex, perpage int64) ([]models.Team, error) {
	return service.repository.Teams(startIndex, perpage)
}

func (service *teamService) TeamsByCountry(country string,startIndex, perpage int64) ([]models.Team, error) {
	return service.repository.TeamsByCountry(country, startIndex, perpage)
}

func lunchUpdateTeamssLoop() {
	tiker := time.NewTicker(time.Hour * 24)
	wg := &sync.WaitGroup{}
	for range tiker.C {
		for _, league := range LocalLeagues {
			teamdto, err := TeamService.footballapi.GetTeamsByLeagueId(league.League_id)
			if err != nil {
				return
			}
			teams := ConvertTeamDtoToTeamModelsObjects(teamdto)
			requiredGoroutines := len(teams)
			wg.Add(requiredGoroutines)
			for _, team := range teams {
				go func(team models.Team, wg *sync.WaitGroup) {
					defer wg.Done()
					TeamService.Upsert(team)
				}(team, wg)
			}

		}
	}

}
