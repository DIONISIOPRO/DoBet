package service

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func CreateBetConsumer(bet_id string) models.BetConsumer {
	consumer := models.BetConsumer{
		BetId: bet_id,
	}
	return consumer
}

func CreateBetProvider(match_id string) models.BetProvider {
	provider := models.BetProvider{
		Match_id:  match_id,
		Consumers: map[string]models.BetConsumer{},
	}
	return provider
}

func ConvertLeagueDtoToLeagueModelObjects(dto dto.LeagueDto) []models.League {
	leagues := []models.League{}
	for i := range dto.Response {
		league := models.League{}
		league.CountryName = dto.Response[i].Country.Name
		league.Logo_url = dto.Response[i].League.Logo
		league.League_id = strconv.FormatInt(dto.Response[i].League.ID, 10)
		league.Name = dto.Response[i].League.Name

		leagues = append(leagues, league)
	}

	return leagues
}

func ConvertTeamDtoToTeamModelsObjects(teamDto dto.TeamDto) []models.Team {
	teams := []models.Team{}
	for i := range teamDto.Response {
		team := models.Team{}
		team.CountryName = teamDto.Response[i].Team.Country
		team.Logo_url = teamDto.Response[i].Team.Logo
		team.CountryName = teamDto.Response[i].Team.Country
		team.Name = teamDto.Response[i].Team.Name
		team.Team_id = strconv.FormatInt(teamDto.Response[i].Team.ID, 10)
		teams = append(teams, team)
	}
	return teams
}

func ConvertMatchDtoToMatchModelsWithoutOddsObjects(matchDto dto.MatchDto) []models.Match {
	matches := []models.Match{}
	for i := range matchDto.Response {
		match := models.Match{}
		matchResult := models.Match_Result{}
		matchResult.All_Scores = matchDto.Response[i].Goals.Home > 0 && matchDto.Response[i].Goals.Away > 0
		matchResult.IsMatchFinished = matchDto.Response[i].Fixture.Status.Short == "FT" || matchDto.Response[i].Fixture.Status.Short == "AET" ||
			matchDto.Response[i].Fixture.Status.Short == "PE"

		if matchDto.Response[i].Teams.Home.Winner {
			matchResult.Winner = "HOME"
		} else if matchDto.Response[i].Teams.Away.Winner {
			matchResult.Winner = "AWAY"
		} else {
			matchResult.Winner = "DRAW"
		}
		matchResult.Team_Away_Goals = int8(matchDto.Response[i].Goals.Away)
		matchResult.Team_Home_Goals = int8(matchDto.Response[i].Goals.Home)
		matchResult.Team_Away_id = strconv.Itoa(int(matchDto.Response[i].Teams.Away.ID))
		matchResult.Team_Home_id = strconv.Itoa(int(matchDto.Response[i].Teams.Home.ID))
		match.Away_team_id = strconv.Itoa(int(matchDto.Response[i].Teams.Away.ID))
		match.Home_team_id = strconv.Itoa(int(matchDto.Response[i].Teams.Home.ID))
		match.Match_id = strconv.Itoa(int(matchDto.Response[i].Fixture.ID))
		match.Result = matchResult
		match.Status = matchDto.Response[i].Fixture.Status.Short
		match.LeagueId = strconv.Itoa(int(matchDto.Response[i].League.ID))
		match.Time = matchDto.Response[i].Fixture.Timestamp
		matches = append(matches, match)
	}
	return matches
}

func ConvertOddDtoToOddModelObjects(oddDto dto.OddsDto) []models.Odds {
	odds := []models.Odds{}
	for _, response := range oddDto.Response {
		odd := models.Odds{}
		for _, ob := range response.Bookmakers[0].Bets {
			switch ob.Name {
			case "Both Teams Score":
				setOddsForBothTeamsScore(odd, ob)
			case "Match Winner":
				setOddsForMatchWinner(odd, ob)
			}
		}
		odds = append(odds, odd)

	}
	return odds

}

func setOddsForBothTeamsScore(odd models.Odds, oddbet dto.OddBet) {
	for _, ov := range oddbet.Values {
		switch ov.Value.(string) {
		case "Yes":
			oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
			odd.AllScoreMarketOdd.Yes = oddIn64
		case "No":
			oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
			odd.AllScoreMarketOdd.No = oddIn64
		}
	}
}

func setOddsForMatchWinner(odd models.Odds, oddbet dto.OddBet) {
	for _, ov := range oddbet.Values {
		switch ov.Value.(string) {
		case "Away":
			oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
			odd.WinnerMarketOdd.Away = oddIn64
		case "Home":
			oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
			odd.WinnerMarketOdd.Home = oddIn64
		case "Draw":
			oddIn64, _ := strconv.ParseFloat(ov.Odd, 64)
			odd.WinnerMarketOdd.Draw = oddIn64
		}
	}
}

func CheckWin(bet models.Bet) bool {
	for _, localbet := range bet.BetGroup {
		if !localbet.IsProcessed {
			return false
		}
	}
	return true
}
