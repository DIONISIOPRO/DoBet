package market

import (
	"github.com/dionisiopro/dobet-bet/domain/option"
	"github.com/dionisiopro/dobet-bet/domain/result"
)

type WinnerMarket struct {
	Will_Team_Away_wins option.Option        `json:"away"`
	Will_Team_Home_wins option.Option               `json:"home"`
	Will_Draw           option.Option               `json:"draw"`
}

type FullWinOption struct {
	WinnerMarket
}

type HalfTimeWinOption struct {
	WinnerMarket
}

func (w WinnerMarket) IsLose(result result.WinResult) bool {
	if w.Will_Team_Away_wins.Selected {
		return !result.AwayWins()
	}
	if w.Will_Team_Home_wins.Selected {
		return !result.HomeWins()
	}
	return !result.IsDraw()
}

func (w WinnerMarket) GetSelectedOdd() float64 {
	if w.Will_Team_Away_wins.Selected {
		return w.Will_Team_Away_wins.Odd
	}
	if w.Will_Team_Home_wins.Selected {
		return w.Will_Team_Home_wins.Odd
	}
	return w.Will_Draw.Odd
}
