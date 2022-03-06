package repositories

import "gitthub.com/dionisiopro/dobet/models"

type BetRepository interface {
	CreateBet(bet models.Bet) (bet_id string ,err error)
	BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error)
	BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error)
	BetById(bet_id string) (models.Bet, error)
	Bets(startIndex, perpage int64) ([]models.Bet, error)
	RunningBets(startIndex, perpage int64) ([]models.Bet, error)
	TotalRunningBetsMoney() float64
	ProcessWin(amount float64, user_id string)
}

type LeagueRepository interface {
	AddLeague(league models.League) error
	DeleteLeague(league_id string) error
	Leagues(startIndex, perpage int64) ([]models.League, error)
}

type MatchRepository interface {
	AddMatch(match models.Match) error
	DeleteMatch(match_id string) error
	UpDateMatch(match_id string, match models.Match) error
	Matches(startIndex, perpage int64) ([]models.Match, error)
	MatchWatch()([]models.Match, error)
}

type TeamRepository interface {
	AddTeam(team models.Team) error
	DeleteTeam(team_id string) error
	Teams(startIndex, perpage int64) ([]models.Team, error)
}

type UserRepository interface {
	Deposit(amount float64, userid string) error
	Withdraw(amount float64, userid string) error
	SignUp(user models.User) error
	Login(user models.User) (models.User, error)
	Users(startIndex, perpage int64) ([]models.User, error)
}