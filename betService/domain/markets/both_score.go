package markets

import (
	"github.com/dionisiopro/dobet-bet/domain/interfaces"
	"github.com/dionisiopro/dobet-bet/domain/option"
)

type BothTimesScoresMarket struct{
	Will_All_Scores     option.Option `json:"allscores"`
}

type BothTimesScoresFullTimeMarket struct {
	BothTimesScoresMarket
}

type BothTimesScoresHalfTimeMarket struct {
	BothTimesScoresMarket
}

func (b BothTimesScoresMarket) IsLose(result interfaces.BothSCoresResult) bool{
	return b.Will_All_Scores.Selected == result.BothScores()
}

func (b BothTimesScoresMarket) GetSelectedOdd() float64{
	return b.Will_All_Scores.Odd
}