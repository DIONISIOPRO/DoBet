package domain

type MatchResultBase struct{
	TeamHomeGoals int
    TeamAwayGoals int
}

type MatchResultImpl struct{
	League_id string
	Match_id string
	WinResultImpl
	BothScoresImpl
}





