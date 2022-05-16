package domain

type WinResultImpl struct{
	MatchResultBase
}

func (w WinResultImpl) IsDraw() bool{
	return w.TeamAwayGoals == w.TeamHomeGoals
}

func (w WinResultImpl) HomeWins() bool{

	return w.TeamHomeGoals > w.TeamAwayGoals
}

func (w WinResultImpl) AwayWins() bool{

	return w.TeamAwayGoals > w.TeamHomeGoals
}