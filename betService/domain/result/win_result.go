package result

type WinResult struct{
	MatchResultBase
}

func (w WinResult) IsDraw() bool{
	return w.TeamAwayGoals == w.TeamHomeGoals
}

func (w WinResult) HomeWins() bool{
	return w.TeamHomeGoals > w.TeamAwayGoals
}

func (w WinResult) AwayWins() bool{
	return w.TeamAwayGoals > w.TeamHomeGoals
}