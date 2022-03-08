package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func TeamDtoToTeamModel(teamDto dto.TeamDto) []models.Team{
	teams := []models.Team{}
	for i, _ := range teamDto.Response {
	team := models.Team{}
	team.CountryName = teamDto.Response[i].Team.Country
	team.Logo_url = teamDto.Response[i].Team.Logo
	team.Name = teamDto.Response[i].Team.Name
	team.Team_id =strconv.FormatInt(teamDto.Response[i].Team.ID, 10) 
	teams = append(teams, team)
	}
	return teams
}