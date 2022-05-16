package models

type Option struct{
	IsSelected bool `json:"is_selected"`
	Odd float64 `json:"odd"`
}


type BothTimesScoresFullTimeOption struct {
	Will_All_Scores     bool `json:"allscores"`
}

type BothTimesScoresHalfTimeOption struct {
	Will_All_Scores     bool `json:"allscores"`
}

type WinOption struct{
	Will_Team_Away_wins Option `json:"away"`
	Will_Team_Home_wins Option `json:"home"`
	Will_Draw        Option `json:"draw"`
}

type HalfTimeOption struct{
	Will_Team_Away_wins Option `json:"away"`
	Will_Team_Home_wins Option `json:"home"`
	Will_Draw           Option `json:"draw"`
}