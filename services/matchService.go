package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func AddMatch(match models.Match) error {
	return repositories.AddMatch(match)
}

func DeleteMatch(match_id string) error {
	return repositories.DeleteMatch(match_id)
}

func UpDateMatch(match_id string, match models.Match) error {
	return repositories.UpDateMatch(match_id, match)
}

func Matches() []models.Match {
	return repositories.Matches()
}