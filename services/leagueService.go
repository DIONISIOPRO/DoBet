package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var leagueRepository repositories.LeagueRepository

func NewLeagueService(leagueRepo repositories.LeagueRepository) {
	leagueRepository = leagueRepo
}

func AddLeague(league models.League) error {
	return leagueRepository.AddLeague(league)
}

func DeleteLeague(league_id string) error {
	return leagueRepository.DeleteLeague(league_id)
}

func Leagues(startIndex, perpage int64) ([]models.League, error) {
	return leagueRepository.Leagues(startIndex, perpage)
}
