package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)


var LeagueService  leagueService
type leagueService struct{
	repo repositories.LeagueRepository
}
func SetupLeagueService(repo repositories.LeagueRepository) *leagueService{
	LeagueService.repo = repo
	return &LeagueService
}

func (service *leagueService)  AddLeague(league models.League) error {
	return service.repo.AddLeague(league)
}

func (service *leagueService) DeleteLeague(league_id string) error {
	return service.repo.DeleteLeague(league_id)
}

func (service *leagueService) Leagues(startIndex, perpage int64) ([]models.League, error) {
	return service.repo.Leagues(startIndex, perpage)
}
