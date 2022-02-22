package repositories

import "gitthub.com/dionisiopro/dobet/models"

func AddLeague(league models.League) error {
	return nil
}

func DeleteLeague(league_id string) error {
	return nil
}

func UpDateLeague(league_id string, league models.League) error {
	return nil
}

func Leagues() []models.League {
	return []models.League{}
}