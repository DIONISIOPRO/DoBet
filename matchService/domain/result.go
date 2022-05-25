package domain

const (
	StatusFT   = "end_ft"
	StatusLive = "live"
	StatusHT   = "end_ht"
)

type MatchResult struct {
	MatchId   string
	Status    string
	HomeGoals int8
	AwayGoals int8
}

func (m MatchResult) AwayWins() bool {
	return m.AwayGoals > m.HomeGoals
}

func (m MatchResult) HomeWins() bool {
	return m.AwayGoals < m.HomeGoals
}

func (m MatchResult) IsDraw() bool {
	return m.AwayGoals == m.HomeGoals
}
