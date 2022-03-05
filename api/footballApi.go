package api

import "gitthub.com/dionisiopro/dobet/models"

type footballapi struct{}

func NewFootBallApi() FootBallApi {
	return &footballapi{}
}

func (f *footballapi) AddLeague(league models.League) error {
	return nil
}

func (f *footballapi) AddMatch(match models.Match) error {
	return nil
}

func (f *footballapi) AddTeam(team models.Team) error {
	return nil
}