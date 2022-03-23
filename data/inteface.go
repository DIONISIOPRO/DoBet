package data

import "gitthub.com/dionisiopro/dobet/models"

type FootballData interface {
	GetLeague(id int64) (models.League, error)
	GetLeagues() ([]models.League, error)
	GetCups() ([]models.League, error)
	GetNext20MatchesByLeagueId(leagueid string) ([]models.Match, error)
	GetLast5MatchesByLeagueId(leagueid string) ([]models.Match, error)
	GetTeamsByLeagueId(league_id string) ([]models.Team, error)
	GetOddsByLeagueId(matchId string) ([]models.Odds, error)
}
