package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func AddLeague(league models.League) error{
	return repositories.AddLeague(league)
}

func DeleteLeague(league_id string) error{
	return repositories.DeleteLeague(league_id)
}

func UpDateLeague(league_id string, league models.League) error{
	return repositories.UpDateLeague(league_id, league)
}

func Leagues() []models.League{
	return repositories.Leagues()
}