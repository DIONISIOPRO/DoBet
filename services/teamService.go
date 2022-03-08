package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)
var TeamService teamService
type teamService struct{
	repo repositories.TeamRepository
}
func SetupTeamSerice(repo repositories.TeamRepository) *teamService {
	TeamService.repo = repo
	return &TeamService
}

func (service *teamService) AddTeam(team models.Team) error {
	return service.repo.AddTeam(team)
}

func (service *teamService) DeleteTeam(team_id string) error {
	return service.repo.DeleteTeam(team_id)
}

func (service *teamService)  Teams(startIndex, perpage int64) ([]models.Team, error) {
	return service.repo.Teams(startIndex, perpage)
}
