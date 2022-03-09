package service

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)


var LeagueService = &leagueService{}
type leagueService struct{
	repo repository.LeagueRepository
}
func SetupLeagueService(leaguerepositorry repository.LeagueRepository){
	LeagueService.repo = leaguerepositorry
	LeagueService.repo = leaguerepositorry
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
