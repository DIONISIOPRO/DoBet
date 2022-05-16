package domain

type MatchResultBase struct{
	League_id string
	Match_id string
	TeamHomeGoals int
    TeamAwayGoals int
}

type MatchResultImpl struct{
	WinResultImpl
	BothScoresImpl
}





