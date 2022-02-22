package repositories

import "gitthub.com/dionisiopro/dobet/models"

func AddTeam(team models.Team) error {
	return nil
}

func DeleteTeam(team_id string) error {
	return nil
}

func UpDateTeam(team_id string, team models.Team) error {
	return nil
}

func Teams() []models.Team {
	return []models.Team{}
}
