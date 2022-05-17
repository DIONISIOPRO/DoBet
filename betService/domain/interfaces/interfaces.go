package domain

type(
	WinResult interface{
		IsDraw() bool
		AwayWins() bool
		HomeWins() bool
	}
	
	BothSCoresResult interface{
		BothScores() bool
	}
	
	MatchResult interface{
		WinResult
		BothSCoresResult
	}
	
	BetMarket interface{
		IsLose(MatchResult) bool
		GetSelectedOdd() float64
	}
	
	BetBase interface{
		IsValid() bool
		IsLose(MatchResult) bool
		GetGlobalOdd()
		GetPotenctialWin() float64
		GetTotalAmount() float64
		IsFinished() bool
	}
	SingleBet interface{
		IsLose(MatchResult) bool
	}
) 