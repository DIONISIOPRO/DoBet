package result


type BothScoresResult struct{
	MatchResultBase
}

func (b BothScoresResult) BothScores() bool{
	return b.TeamAwayGoals > 0 && b.TeamHomeGoals > 0
}