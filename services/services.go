package services

import "gitthub.com/dionisiopro/dobet/models"

type BetService interface {
	CreateBet(owner_id string, match_id []string, odd float64, amount float64, market []models.Market, options []models.BetOption) error
	BetByUser(user_id string) ([]models.Bet, error)
	BetByMatch(match_id string) ([]models.Bet, error)
	Bets() ([]models.Bet, error)
	RunningBets() ([]models.Bet, error)
	TotalRunningBetsMoney() float32
	ProcessBet(bet_id string, match_result models.Match_Result)
	ProcessWin(amount float64, user_id string)
	CreateBetConsumer(bet_id string) BetConsumer
	CreateBetProvider(match_id string) BetProvider
}

type LeagueService interface{
	AddLeague(league models.League) error
	DeleteLeague(league_id string) error
	UpDateLeague(league_id string, league models.League) error
	Leagues() ([]models.League, error)
}

type MatchService interface{
	AddMatch(match models.Match) error
	DeleteMatch(match_id string) error
	UpDateMatch(match_id string, match models.Match) error
	Matches() ([]models.Match, error)
}

type TeamService interface{
	AddTeam(team models.Team) error
	DeleteTeam(team_id string) error
	UpDateTeam(team_id string, team models.Team) error
	Teams() ([]models.Team, error)
}

type UserService interface{
	Deposit(amount float64, userid string) error
	Withdraw(amount float64, userid string) error
	Login(user models.User) error
	SignUp(user models.User) error
	Users() ([]models.User, error)
}