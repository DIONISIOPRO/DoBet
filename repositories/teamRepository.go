package repositories

import (
	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
)

var teamCollection = database.OpenCollection("teams")

type teamRepository struct{

}

func NewTeamRepository() TeamRepository{
	return &teamRepository{}
}


func (repo *teamRepository) AddTeam(team models.Team) error {
	return nil
}

func (repo *teamRepository)  DeleteTeam(team_id string) error {
	return nil
}

func(repo *teamRepository)  UpDateTeam(team_id string, team models.Team) error {
	return nil
}

func(repo *teamRepository)  Teams() ([]models.Team, error ){
	return []models.Team{}, nil
}
