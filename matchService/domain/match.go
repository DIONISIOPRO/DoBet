package domain

import "time"

type Match struct {
	Id string
	LeagueId  string
	StartTime time.Time
	EndTime   time.Time
	Home Team
	Away Team
}
