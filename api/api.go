package api

import "gitthub.com/dionisiopro/dobet/models"

type FootBallApi interface {
	AddLeague(league models.League) error
	AddTeam(team models.Team) error
	AddMatch(match models.Match) error
	UpDateMatch(match_id string, match models.Match) error 
}

type PaymentApi interface {
	Deposit(amount float64, userid string) error
	Withdraw(amount float64, userid string) error
} 