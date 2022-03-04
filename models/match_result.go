package models

type Match_Result struct {
	Match_id          string
	Team_Home_id      string
	Team_Away_id      string
	Team_Away_Goals   int8
	Team_Home_Goals   int8
	Is_Team_Home_wins bool
	Is_Team_Away_wins bool
	Is_Draw           bool
	All_Scores        bool
	IsMatchFinished   bool
}
