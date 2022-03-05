package repositories

import (
	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
)

var leagueCollection = database.OpenCollection("leagues")

type leagueRepository struct{}

func NewLeagueRepository() LeagueRepository{
	return &leagueRepository{}
}

func (service *leagueRepository) AddLeague(league models.League) error {
	return nil
}

func(service *leagueRepository)  DeleteLeague(league_id string) error {
	return nil
}

func(service *leagueRepository)  UpDateLeague(league_id string, league models.League) error {
	return nil
}

func (service *leagueRepository) Leagues() ([]models.League, error) {
	return []models.League{}, nil
}