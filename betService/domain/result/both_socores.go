package domain


type BothScoresResultImpl struct{
	MatchResultBase
}

func (b BothScoresResultImpl) BothScores() bool{
	return b.TeamAwayGoals > 0 && b.TeamHomeGoals > 0
}