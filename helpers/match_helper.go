package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/dto"

)

func MatchDtoRespoonseToMatchModel(matchDto dto.MatchDto) (models.Match, error) {
	match := models.Match{}
	odds := models.Odds{}
	matchResult := models.Match_Result{}
	matchResult.All_Scores = matchDto.Response[0].Goals.Home > 0 && matchDto.Response[0].Goals.Away > 0
	matchResult.IsMatchFinished = matchDto.Response[0].Fixture.Status.Short == "FT" || matchDto.Response[0].Fixture.Status.Short == "AET" ||
		matchDto.Response[0].Fixture.Status.Short == "PE"
	matchResult.Is_Draw = matchDto.Response[0].Goals.Home == matchDto.Response[0].Goals.Away
	matchResult.Is_Team_Away_wins = matchDto.Response[0].Teams.Away.Winner
	matchResult.Is_Team_Home_wins = matchDto.Response[0].Teams.Home.Winner
	matchResult.Team_Away_Goals = int8(matchDto.Response[0].Goals.Away)
	matchResult.Team_Home_Goals = int8(matchDto.Response[0].Goals.Home)
	matchResult.Team_Away_id = strconv.Itoa(int(matchDto.Response[0].Teams.Away.ID))
	matchResult.Team_Home_id = strconv.Itoa(int(matchDto.Response[0].Teams.Home.ID))
	match.Away_team_id = strconv.Itoa(int(matchDto.Response[0].Teams.Away.ID))
	match.Home_team_id = strconv.Itoa(int(matchDto.Response[0].Teams.Home.ID))
	match.Match_id = strconv.Itoa(int(matchDto.Response[0].Fixture.ID))
	match.Result = matchResult
	match.Status = matchDto.Response[0].Fixture.Status.Short

	matchId, _ := strconv.Atoi(match.Match_id)
	oddDto, err := api.GetOddsByMatchId(matchId)
	if err != nil {
		return models.Match{}, nil
	}
	odds = OddDtoToOddModel(oddDto)
	match.Odds = odds
	return match, nil
}