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
	GetLeague(id int64) (dto.LeagueDto, error)
	GetLeagues() (dto.LeagueDto, error)
	GetCups() (dto.LeagueDto, error)
	GetNext20MatchesByLeagueId(leagueid string) (dto.MatchDto, error)
	GetLast5MatchesByLeagueId(leagueid string) (dto.MatchDto, error)
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
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/leagues?type=league&season=%v", api.BaseUrl, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.LeagueDto{}, nil
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.LeagueDto{}, err
	}
	defer response.Body.Close()

	var leagues = dto.LeagueDto{}

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

func (api footballapi) GetLeague(id int64) (dto.LeagueDto, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/leagues?id=%vtype=league&season=%v", api.BaseUrl, id, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.LeagueDto{}, nil
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.LeagueDto{}, err
	}
	defer response.Body.Close()

	var leagues = dto.LeagueDto{}

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
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/leagues?type=cup&season=%v", api.BaseUrl, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.LeagueDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.LeagueDto{}, err
	}
	defer response.Body.Close()

	var leagues = dto.LeagueDto{}

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

func (api footballapi) GetNext20MatchesByLeagueId(leagueid string) (dto.MatchDto, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/fixtures?league=%v&season=%v&next=20", api.BaseUrl, leagueid, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.MatchDto{}, err
	}
	defer response.Body.Close()

	var matches = dto.MatchDto{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.MatchDto{}, err
	}

	if err = json.Unmarshal(data, &matches); err != nil {
		return dto.MatchDto{}, err
	}

	return matches, nil
}

func (api footballapi) GetLast5MatchesByLeagueId(leagueid string) (dto.MatchDto, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/fixtures?league=%v&season=%v&last=5", api.BaseUrl, leagueid, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.MatchDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.MatchDto{}, err
	}
	defer response.Body.Close()

	var matches = dto.MatchDto{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.MatchDto{}, err
	}

	if err = json.Unmarshal(data, &matches); err != nil {
		return dto.MatchDto{}, err
	}

	return matches, nil
}

func (api footballapi) GetTeamsByLeagueId(league string) (dto.TeamDto, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/teams?league=%v&season=%v", api.BaseUrl, league, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.TeamDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.TeamDto{}, err
	}
	defer response.Body.Close()

	var teams = dto.TeamDto{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.TeamDto{}, err
	}
	err = json.Unmarshal(data, &teams)
	if err != nil {
		return dto.TeamDto{}, err
	}
	return teams, nil
}

func (api footballapi) GetOddsByLeagueId(leagueid string) (dto.OddsDto, error) {
	//	season := time.Now().Year() - 1
	url1 := "https://v3.football.api-sports.io/odds?page=1&league=39&season=2021"
	//	url := fmt.Sprintf("%v/odds?page=%v&league=%v&season=%v",api.BaseUrl,1, leagueid, season)
	var req, err = http.NewRequest("GET", url1, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return dto.OddsDto{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return dto.OddsDto{}, err
	}
	defer response.Body.Close()

	var odd = dto.OddsDto{}

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
