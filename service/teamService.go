package service

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)
var TeamService = &teamService{}
type teamService struct{
	repository repository.TeamRepository
}
func SetupTeamSerice(teamRepository repository.TeamRepository)  {
	TeamService.repository = teamRepository
}

func (service *teamService) AddTeam(team models.Team) error {
	return service.repository.AddTeam(team)
}

func (service *teamService) DeleteTeam(team_id string) error {
	return service.repository.DeleteTeam(team_id)
}

func (service *teamService)  Teams(startIndex, perpage int64) ([]models.Team, error) {
	return service.repository.Teams(startIndex, perpage)
}
