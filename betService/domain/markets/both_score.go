package domain

type BothTimesScoresMarket struct{
	result BothSCoresResult `json:"result"`
	Will_All_Scores     Option `json:"allscores"`
}
type BothTimesScoresFullTimeMarket struct {
	BothTimesScores
}

type BothTimesScoresHalfTimeMarket struct {
	BothTimesScores
}

func (b BothTimesScores) IsLose() bool{
	return b.Will_All_Scores.Selected == b.result.BothScores()
}

func (b BothTimesScores) GetSelectedOdd() float64{
	return b.Will_All_Scores.Odd
}