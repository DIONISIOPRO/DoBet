package domain

type BothTimesScoresMarket struct{
	Will_All_Scores     Option `json:"allscores"`
}

type BothTimesScoresFullTimeMarket struct {
	BothTimesScores
}

type BothTimesScoresHalfTimeMarket struct {
	BothTimesScores
}

func (b BothTimesScores) IsLose(result BothSCoresResult) bool{
	return b.Will_All_Scores.Selected == result.BothScores()
}

func (b BothTimesScores) GetSelectedOdd() float64{
	return b.Will_All_Scores.Odd
}