package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var teamRepository repositories.TeamRepository

func NewTeamSerice(repo repositories.TeamRepository) {
	teamRepository = repo
}

func AddTeam(team models.Team) error {
	return teamRepository.AddTeam(team)
}

func DeleteTeam(team_id string) error {
	return teamRepository.DeleteTeam(team_id)
}

func Teams(startIndex, perpage int64) ([]models.Team, error) {
	return teamRepository.Teams(startIndex, perpage)
}
