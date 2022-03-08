package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func MatchDtoRespoonseToMatchModelWithoutOdds(matchDto dto.MatchDto) ([]models.Match, error) {
	matches := []models.Match{}
	for i, _ := range matchDto.Response {
		match := models.Match{}
		matchResult := models.Match_Result{}
		matchResult.All_Scores = matchDto.Response[i].Goals.Home > 0 && matchDto.Response[i].Goals.Away > 0
		matchResult.IsMatchFinished = matchDto.Response[i].Fixture.Status.Short == "FT" || matchDto.Response[i].Fixture.Status.Short == "AET" ||
			matchDto.Response[i].Fixture.Status.Short == "PE"
		matchResult.Is_Draw = matchDto.Response[i].Goals.Home == matchDto.Response[i].Goals.Away
		matchResult.Is_Team_Away_wins = matchDto.Response[i].Teams.Away.Winner
		matchResult.Is_Team_Home_wins = matchDto.Response[i].Teams.Home.Winner
		matchResult.Team_Away_Goals = int8(matchDto.Response[i].Goals.Away)
		matchResult.Team_Home_Goals = int8(matchDto.Response[i].Goals.Home)
		matchResult.Team_Away_id = strconv.Itoa(int(matchDto.Response[i].Teams.Away.ID))
		matchResult.Team_Home_id = strconv.Itoa(int(matchDto.Response[i].Teams.Home.ID))
		match.Away_team_id = strconv.Itoa(int(matchDto.Response[i].Teams.Away.ID))
		match.Home_team_id = strconv.Itoa(int(matchDto.Response[i].Teams.Home.ID))
		match.Match_id = strconv.Itoa(int(matchDto.Response[i].Fixture.ID))
		match.Result = matchResult
		match.Status = matchDto.Response[i].Fixture.Status.Short
		//TODO
	//	match.Time = matchDto.Response[i].Fixture.Timestamp
		matches = append(matches, match)
	}
	return matches, nil
}
