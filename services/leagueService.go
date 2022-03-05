package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

type leagueServce struct{
	leagueRepo repositories.LeagueRepository
}

func NewLeagueService(leagueRepo repositories.LeagueRepository) LeagueService{
 return &leagueServce{
	 leagueRepo: leagueRepo,
 }
}

func (service *leagueServce) AddLeague(league models.League) error{
	return service.leagueRepo.AddLeague(league)
}

func (service *leagueServce) DeleteLeague(league_id string) error{
	return service.leagueRepo.DeleteLeague(league_id)
}

func (service *leagueServce) UpDateLeague(league_id string, league models.League) error{
	return service.leagueRepo.UpDateLeague(league_id,league)
}

func (service *leagueServce) Leagues(startIndex, perpage int64) ([]models.League, error){
	return service.leagueRepo.Leagues(startIndex, perpage)
}