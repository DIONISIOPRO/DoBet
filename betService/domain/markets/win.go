package domain

type WinOption struct{
	result WinResult `json:"result"`
	Will_Team_Away_wins Option `json:"away"`
	Will_Team_Home_wins Option `json:"home"`
	Will_Draw        Option `json:"draw"`
}

type FullWinOption struct{
	WinOption
}

type HalfTimeWinOption struct{
	WinOption
}

func (w WinOption) IsLose(result WinResult) bool{
	if w.Will_Team_Away_wins.Selected{
		return !result.AwayWins()
	}
	if w.Will_Team_Home_wins.Selected{
		return !result.HomeWins()
	}
	return !result.IsDraw()
}

func (w WinOption) GetSelectedOdd() float64{
	if w.Will_Team_Away_wins.Selected{
		return w.Will_Team_Away_wins.Odd
	}
	if w.Will_Team_Home_wins.Selected{
		return w.Will_Team_Home_wins.Odd
	}
	return Will_Draw.Odd
}

