package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gitthub.com/dionisiopro/dobet/dto"
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

type FootBallApi interface {
	GetLeagues() (dto.LeagueDto, error)
	GetCups() (dto.LeagueDto, error)
	GetMatchesByLeagueId(leagueid string) (dto.MatchDto, error)
	GetMatchesByCupId(leagueid string) (dto.MatchDto, error)
	GetTeamsByLeagueId(league_id string) (dto.TeamDto, error)
	GetOddsByLeagueId(matchId string) (dto.OddsDto, error)
}

func NewFootBallApi(client *http.Client, baseUrl, token, host string) FootBallApi {
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

func (api footballapi) GetLeagues() (dto.LeagueDto, error) {
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

func (api footballapi) GetCups() (dto.LeagueDto, error) {
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

func (api footballapi) GetMatchesByLeagueId(leagueid string) (dto.MatchDto, error) {
	season := time.Now().Year()
	url := fmt.Sprintf("%vleague=%v&season=%v&next=%v",api.BaseUrl, leagueid,season,15)
	var req, err = http.NewRequest("GET", url, nil)
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

func (api footballapi) GetMatchesByCupId(cupid string) (dto.MatchDto, error) {
	season := time.Now().Year()
	url := fmt.Sprintf("%vcup=%v&season=%v&next=%v",api.BaseUrl, cupid,season,15)
	var req, err = http.NewRequest("GET", url, nil)
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

func (api footballapi) GetTeamsByLeagueId(teamId string) (dto.TeamDto, error) {
	season := time.Now().Year()
	url := fmt.Sprintf("%v?league=%v&season=%v",api.BaseUrl, teamId,season)
	var req, err = http.NewRequest("GET", url, nil)
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


func (api footballapi) GetOddsByLeagueId(leagueid string) (dto.OddsDto, error) {
	season := time.Now().Year()
	url := fmt.Sprintf("%vleague=%v&season=%v&page=%v",api.BaseUrl, leagueid,season,1)
	var req, err = http.NewRequest("GET", url, nil)
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
