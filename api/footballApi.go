package api

import (
	"net/http"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

type footballapi struct{
	BaseUrl string
	Client *http.Client
	Header Header
}


type Header struct{
	Token string
	Host string

}

func NewFootBallApi(client *http.Client, baseUrl, token, host string) footballapi{
	header := Header{
		Token: token,
		Host: host,
	}
	api := footballapi{
		BaseUrl: baseUrl,
		Client: client,
		Header: header,
		
	}
	return api
}


func (api *footballapi) GetLeagues() ([]dto.LeagueDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return []dto.LeagueDto{}, nil
	}

	response, err := api.Client.Do(req)
	defer response.Body.Close()

	return []dto.LeagueDto{}, nil
}

func(api *footballapi) GetCups() ([]dto.LeagueDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return []dto.LeagueDto{}, nil
	}
	response, err := api.Client.Do(req)
	defer response.Body.Close()
	return []dto.LeagueDto{}, nil
}

func (api *footballapi)GetMatchesByLeagueId(leagueid string) ([]dto.MatchDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return []dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	defer response.Body.Close()
	return []dto.MatchDto{}, err
}

func (api *footballapi) GetMatchesByCupId(leagueid string) ([]dto.MatchDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return []dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	defer response.Body.Close()
	return []dto.MatchDto{}, err
}

func(api *footballapi) GetTeamsByLeagueId(team models.Team) ([]dto.TeamDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return []dto.TeamDto{}, err
	}
	response, err := api.Client.Do(req)
	defer response.Body.Close()
	return []dto.TeamDto{}, err
}

func(api *footballapi) GetMatchByiD(match_id string, match models.Match) (dto.MatchDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	defer response.Body.Close()
	return dto.MatchDto{}, err
}

func(api *footballapi) GetOddsByMatchId(matchId int) (dto.OddsDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.OddsDto{}, err
	}
	response, err := api.Client.Do(req)
	defer response.Body.Close()
	return dto.OddsDto{}, nil
}
