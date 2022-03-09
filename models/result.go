package models

type Match_Result struct {
	Match_id        string
	Team_Home_id    string
	Team_Away_id    string
	Team_Away_Goals int8
	Team_Home_Goals int8
	Winner          string
	All_Scores      bool
	IsMatchFinished bool
}
