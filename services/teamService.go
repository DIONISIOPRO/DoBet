package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func AddTeam(team models.Team) error {
	return repositories.AddTeam(team)
}

func DeleteTeam(team_id string) error {
	return repositories.DeleteTeam(team_id)
}

func UpDateTeam(team_id string, team models.Team) error {
	return repositories.UpDateTeam(team_id, team)
}

func Teams() []models.Team{
	return repositories.Teams()
}
