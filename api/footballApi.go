package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
)

type footballapi struct {
	BaseUrl string
	Client  *http.Client
	Header  Header
}

type Header struct {
	Token string
	Host  string
}

type FootBallApi interface{
	GetLeagues() (dto.LeagueDto, error)
	GetCups() (dto.LeagueDto, error)
	GetMatchesByLeagueId(leagueid int) (dto.MatchDto, error)
	GetMatchesByCupId(leagueid int) (dto.MatchDto, error)
	GetTeamsByLeagueId(league_id int) (dto.TeamDto, error)
	GetMatchByiD(match_id string) (dto.MatchDto, error)
	GetOddsByMatchId(matchId int) (dto.OddsDto, error)
	Matches() (dto.MatchDto, error)
}

func NewFootBallApi(client *http.Client, baseUrl, token, host string) footballapi {
	header := Header{
		Token: token,
		Host:  host,
	}
	api := footballapi{
		BaseUrl: baseUrl,
		Client:  client,
		Header:  header,
	}
	return api
}

func (api *footballapi) GetLeagues() (dto.LeagueDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.LeagueDto{}, nil
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.LeagueDto{}, nil
	}
	defer response.Body.Close()

	var leagues dto.LeagueDto

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return dto.LeagueDto{}, nil
	}

	err = json.Unmarshal(data, &leagues)
	if err != nil {
		return dto.LeagueDto{}, nil
	}

	return leagues, nil
}

func (api *footballapi) GetCups() (dto.LeagueDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.LeagueDto{}, nil
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.LeagueDto{}, nil
	}
	defer response.Body.Close()

	var leagues dto.LeagueDto

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return dto.LeagueDto{}, nil
	}

	err = json.Unmarshal(data, &leagues)
	if err != nil {
		return dto.LeagueDto{}, nil
	}

	return leagues, nil
}

func (api *footballapi) GetMatchesByLeagueId(leagueid string) (dto.MatchDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.MatchDto{}, err
	}
	defer response.Body.Close()

	var matches dto.MatchDto

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.MatchDto{}, err
	}

	if err = json.Unmarshal(data, &matches); err != nil {
		return dto.MatchDto{}, err
	}

	return matches, nil
}

func (api *footballapi) GetMatchesByCupId(leagueid string) (dto.MatchDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.MatchDto{}, err
	}
	defer response.Body.Close()

	if err != nil {
		return dto.MatchDto{}, err
	}
	defer response.Body.Close()

	var matches dto.MatchDto

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.MatchDto{}, err
	}

	if err = json.Unmarshal(data, &matches); err != nil {
		return dto.MatchDto{}, err
	}

	return matches, nil
}

func (api *footballapi) GetTeamsByLeagueId(team models.Team) (dto.TeamDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.TeamDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.TeamDto{}, err
	}
	defer response.Body.Close()

	var teams dto.TeamDto

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.TeamDto{}, err
	}
	json.Unmarshal(data, &teams)

	return teams, nil
}

func (api *footballapi) GetMatchByiD(match_id string) (dto.MatchDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.MatchDto{}, err
	}
	defer response.Body.Close()

	var match dto.MatchDto

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return dto.MatchDto{}, err
	}

	if err = json.Unmarshal(data, &match); err != nil {
		return dto.MatchDto{}, err
	}
	return match, err
}

func (api *footballapi) GetOddsByMatchId(matchId int) (dto.OddsDto, error) {
	var req, err = http.NewRequest("GET", api.BaseUrl, nil)
	if err != nil {
		return dto.OddsDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.OddsDto{}, err
	}
	defer response.Body.Close()

	var odd dto.OddsDto

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.OddsDto{}, err
	}

	err = json.Unmarshal(data, &odd)
	if err != nil {
		return dto.OddsDto{}, err
	}

	return odd, nil
}
