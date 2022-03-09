package models

type Odds struct {
	Match_id                 string
	All_Teams_Scores_odd     float32
	Not_All_Teams_Scores_odd float32
	Team_Away_Win_odd        float32
	Team_Home_Win_odd        float32
	Draw_odd                 float32
}
