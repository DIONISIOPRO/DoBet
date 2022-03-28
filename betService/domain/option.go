package models

type BetOption struct {
	Will_All_Scores     bool `json:"allscores"`
	Will_Team_Away_wins bool `json:"away"`
	Will_Team_Home_wins bool `json:"home"`
	Will_Draw           bool `json:"draw"`
}
