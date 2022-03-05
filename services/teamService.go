package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

type teamService struct{
	repo repositories.TeamRepository
}
func NewTeamSerice(repo repositories.TeamRepository) TeamService{
	return &teamService{
		repo: repo,
	}

}

func (service *teamService) AddTeam(team models.Team) error {
	return service.repo.AddTeam(team)
}

func (service *teamService)  DeleteTeam(team_id string) error {
	return service.repo.DeleteTeam(team_id)
}

func (service *teamService)  UpDateTeam(team_id string, team models.Team) error {
	return service.repo.UpDateTeam(team_id, team)
}

func (service *teamService)  Teams(startIndex, perpage int64) ([]models.Team, error){
	return service.repo.Teams(startIndex, perpage)
}
