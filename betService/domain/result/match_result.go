package result

type MatchResultBase struct{
	TeamHomeGoals int
    TeamAwayGoals int
}

type MatchResult struct{
	League_id string
	Match_id string
	WinResult
	BothScoresResult
}





