package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func LeagueDtoToLeagueModel(dto dto.LeagueDto) models.League{
	league := models.League{}
	league.CountryName = dto.Response[0].Country.Name
	league.Logo_url = dto.Response[0].League.Logo
	league.League_id = strconv.FormatInt(dto.Response[0].League.ID, 10)
	league.Name = dto.Response[0].League.Name

	return league
}