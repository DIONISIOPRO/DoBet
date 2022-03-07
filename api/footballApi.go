package api

import (
	"net/http"

	"gitthub.com/dionisiopro/dobet/models"
)

const(
	url = ""
)
var Client = &http.Client{}

type footballapi struct{}


func  GetLeagues() error {
	var _, _ = http.NewRequest("GET", url, nil)

	return nil
}

func  GetCups() error {
	var _, _ = http.NewRequest("GET", url, nil)

	return nil
}

func  GetMatchesByLeagueId(leagueid string) error {
	var _, _ = http.NewRequest("GET", url, nil)

	return nil
}

func GetMatchesByCupId(leagueid string) error {
	var _, _ = http.NewRequest("GET", url, nil)

	return nil
}

func GetTeamsByLeagueId(team models.Team) error {
	var _, _ = http.NewRequest("GET", url, nil)

	return nil
}

func  GetMatchByiD(match_id string, match models.Match) error{
	var _, _ = http.NewRequest("GET", url, nil)

	return nil
}