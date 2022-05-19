package market

import "github.com/dionisiopro/dobet-bet/domain/result"

type MatchMarket struct{
	Match_id string
	BothTimesScoresMarket
	WinnerMarket
}

func (m MatchMarket) GetSelectedOdd() float64{
	return m.GetBothScoresSelectedOdd() + m.GetWinnerSelectedOdd()
}

func (m MatchMarket) IsLose(result result.MatchResult) bool{
	return m.BothSCoresMarketIsLose(result) || m.WinnerMarketIsLose(result)
}