package repositories

import "gitthub.com/dionisiopro/dobet/models"

type BetRepository interface {
	CreateBet(bet models.Bet) (bet_id string ,err error)
	BetByUser(user_id string) ([]models.Bet, error)
	BetByMatch(match_id string) ([]models.Bet, error)
	BetById(bet_id string) (models.Bet, error)
	Bets() ([]models.Bet, error)
	RunningBets() ([]models.Bet, error)
	TotalRunningBetsMoney() float32
	BetWatch()
	ProcessWin(amount float64, user_id string)
}

type LeagueRepository interface {
	AddLeague(league models.League) error
	DeleteLeague(league_id string) error
	UpDateLeague(league_id string, league models.League) error
	Leagues() ([]models.League, error)
}

type MatchRepository interface {
	AddMatch(match models.Match) error
	DeleteMatch(match_id string) error
	UpDateMatch(match_id string, match models.Match) error
	Matches() ([]models.Match, error)
}

type TeamRepository interface {
	AddTeam(team models.Team) error
	DeleteTeam(team_id string) error
	UpDateTeam(team_id string, team models.Team) error
	Teams() ([]models.Team, error)
}

type UserRepository interface {
	Deposit(amount float64, userid string) error
	Withdraw(amount float64, userid string) error
	Login(user models.User) error
	SignUp(user models.User) error
	Users() ([]models.User, error)
	Win(amount float64, user_id string)
}