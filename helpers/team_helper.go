package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func TeamDtoToTeamModel(teamDto dto.TimeDto) models.Team{
	team := models.Team{}
	team.CountryName = teamDto.Response[0].Team.Country
	team.Logo_url = teamDto.Response[0].Team.Logo
	team.Name = teamDto.Response[0].Team.Name
	team.Team_id =strconv.FormatInt(teamDto.Response[0].Team.ID, 10) 
	return team
}