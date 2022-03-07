package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
)

func CustomBet(
	owner_id string,
	match_id []string,
	odd float64,
	amount float64,
	market []models.Market,
	options []models.BetOption,
	) models.Bet {
		
	bet := models.Bet{
		Bet_owner:     owner_id,
		Match_id:      match_id,
		Amount:        amount,
		Markets:       market,
		Options:       options,
		Odd:           odd,
		Potencial_win: odd * amount,
		IsProcessed:   false,
		IsFinished:    false,
	}
	return bet
}


func MatchDtoRespoonseToMatchModel(matchDto api.MatchDto) models.Match{
	match := models.Match{}
//	odds := models.Odds{}
	matchResult := models.Match_Result{}
	matchResult.All_Scores = matchDto.Response[0].Goals.Away > 0 && matchDto.Response[0].Goals.Away > 0
	matchResult.IsMatchFinished = matchDto.Response[0].Fixture.Status.Short == "FT" || matchDto.Response[0].Fixture.Status.Short == "AET" ||
	matchDto.Response[0].Fixture.Status.Short == "PE"
	matchResult.Is_Draw = matchDto.Response[0].Goals.Away ==  matchDto.Response[0].Goals.Away 
	matchResult.Is_Team_Away_wins = matchDto.Response[0].Teams.Away.Winner
	matchResult.Is_Team_Home_wins = matchDto.Response[0].Teams.Home.Winner
	matchResult.Team_Away_Goals = int8(matchDto.Response[0].Goals.Away)
	matchResult.Team_Home_Goals = int8(matchDto.Response[0].Goals.Home)
	matchResult.Team_Away_id = strconv.Itoa(int(matchDto.Response[0].Teams.Away.ID))
	matchResult.Team_Home_id = strconv.Itoa(int(matchDto.Response[0].Teams.Home.ID))

	//odds.All_Teams_Scores_odd = matchDto.Response[0].
	
	match.Away_team_id = strconv.Itoa(int(matchDto.Response[0].Teams.Away.ID))
	match.Home_team_id = strconv.Itoa(int(matchDto.Response[0].Teams.Home.ID))
	match.Match_id = strconv.Itoa(int(matchDto.Response[0].Fixture.ID))
	match.Result = matchResult
	match.Status = matchDto.Response[0].Fixture.Status.Short

	return match
}