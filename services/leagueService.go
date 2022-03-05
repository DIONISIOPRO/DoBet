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

func (l *leagueServce) AddLeague(league models.League) error{
	return l.leagueRepo.AddLeague(league)
}

func (l *leagueServce) DeleteLeague(league_id string) error{
	return l.leagueRepo.DeleteLeague(league_id)
}

func (l *leagueServce) UpDateLeague(league_id string, league models.League) error{
	return l.leagueRepo.UpDateLeague(league_id,league)
}

func (l *leagueServce) Leagues() ([]models.League, error){
	return l.leagueRepo.Leagues()
}