package helpers

import (
	"strconv"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

func LeagueDtoToLeagueModel(dto dto.LeagueDto) []models.League {
	leagues := []models.League{}
	for i, _ := range dto.Response {
		league := models.League{}
		league.CountryName = dto.Response[i].Country.Name
		league.Logo_url = dto.Response[i].League.Logo
		league.League_id = strconv.FormatInt(dto.Response[i].League.ID, 10)
		league.Name = dto.Response[i].League.Name

		leagues = append(leagues, league)
	}

	return leagues
}
