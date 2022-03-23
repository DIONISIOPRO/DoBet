package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitthub.com/dionisiopro/dobet/data"
	"gitthub.com/dionisiopro/dobet/dto"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/service"
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

func NewFootBallApi(client *http.Client, baseUrl, token, host string) data.FootballData {
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

func (api footballapi) GetLeagues() ([]models.League, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/leagues?type=league&season=%v", api.BaseUrl, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return nil, nil
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var leagues = []models.League{}
	var leaguedto = dto.LeagueDto{}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, nil
	}

	err = json.Unmarshal(data, &leaguedto)
	if err != nil {
		return nil, nil
	}
	leagues = service.ConvertLeagueDtoToLeagueModelObjects(leaguedto)

	return leagues, nil
}

func (api footballapi) GetLeague(id int64) (models.League, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/leagues?id=%vtype=league&season=%v", api.BaseUrl, id, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return models.League{}, nil
	}
	response, err := api.Client.Do(req)
	log.Printf("Response: %v", response.Body)
	if err != nil {
		log.Printf("Erorr: %v", err.Error())
		return models.League{}, err
	}
	defer response.Body.Close()

	var leagues = []models.League{}
	var leaguedto = dto.LeagueDto{}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return models.League{}, nil
	}

	err = json.Unmarshal(data, &leaguedto)
	if err != nil {
		return models.League{}, nil
	}
	leagues = service.ConvertLeagueDtoToLeagueModelObjects(leaguedto)
if len(leagues) == 0{
	return models.League{}, nil
}
	return leagues[0], nil
}

func (api footballapi) GetCups() ([]models.League, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/leagues?type=cup&season=%v", api.BaseUrl, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return nil, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var leagues = []models.League{}
	var leaguedto = dto.LeagueDto{}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, nil
	}

	err = json.Unmarshal(data, &leaguedto)
	if err != nil {
		return nil, nil
	}
	leagues = service.ConvertLeagueDtoToLeagueModelObjects(leaguedto)

	return leagues, nil
}

func (api footballapi) GetNext20MatchesByLeagueId(leagueid string) ([]models.Match, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/fixtures?league=%v&season=%v&next=20", api.BaseUrl, leagueid, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return nil, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var matches = []models.Match{}
	var matchesdto = dto.MatchDto{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &matchesdto); err != nil {
		return nil, err
	}
	matches = service.ConvertMatchDtoToMatchModelsWithoutOddsObjects(matchesdto)

	return matches, nil
}

func (api footballapi) GetLast5MatchesByLeagueId(leagueid string) ([]models.Match, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/fixtures?league=%v&season=%v&last=5", api.BaseUrl, leagueid, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return nil, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var matches = []models.Match{}
	var matchesdto = dto.MatchDto{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &matchesdto); err != nil {
		return nil, err
	}
	matches = service.ConvertMatchDtoToMatchModelsWithoutOddsObjects(matchesdto)

	return matches, nil
}

func (api footballapi) GetTeamsByLeagueId(league_id string) ([]models.Team, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/teams?league=%v&season=%v", api.BaseUrl, league_id, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return nil, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var teamsDto = dto.TeamDto{}
	var teams = []models.Team{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &teamsDto)
	if err != nil {
		return nil, err
	}
	teams = service.ConvertTeamDtoToTeamModelsObjects(teamsDto)
	return teams, nil
}

func (api footballapi) GetOddsByLeagueId(matchId string) ([]models.Odds, error) {
	season := time.Now().Year() - 1
	url := fmt.Sprintf("%v/odds?page=%v&league=%v&season=%v", api.BaseUrl, 1, season)
	var req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("x-rapidapi-host", api.Header.Host)
	req.Header.Set("x-apisports-key", api.Header.Token)
	req.Close = true
	if err != nil {
		return []models.Odds{}, err
	}
	response, err := api.Client.Do(req)
	if err != nil {
		return []models.Odds{}, err
	}
	defer response.Body.Close()

	var odddto = dto.OddsDto{}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []models.Odds{}, err
	}

	err = json.Unmarshal(data, &odddto)
	odd := service.ConvertOddDtoToOddModelObjects(odddto)
	if err != nil {
		return []models.Odds{}, err
	}

	return odd, nil
}
