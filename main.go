package main

import (
	"encoding/json"
	"net/http"
	"os"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/config"
	"gitthub.com/dionisiopro/dobet/repository"
	"gitthub.com/dionisiopro/dobet/service"
)

func main() {
	Config := config.BaseConfig{}

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileDecoder := json.NewDecoder(file)
	err = fileDecoder.Decode(&Config)
	if err != nil {
		panic(err)
	}

	const (
		USERCOLLECTION   = "users"
		ODDCOLLECTION    = "odds"
		LEAGUECOLLECTION = "leagues"
		MATCHCOLLECTION  = "matches"
		TEAMCOLLECTION   = "teams"
		BETCOLLECTION    = "bets"
	)


	userRepository := repository.NewUserRepository(USERCOLLECTION)
	oddRepository := repository.NewOddRepository(ODDCOLLECTION)
	leagueRepository := repository.NewLeagueRepository(LEAGUECOLLECTION)
	matchRepository := repository.NewMatchReposiotry(MATCHCOLLECTION)
	teamRepository := repository.NewTeamRepository(TEAMCOLLECTION)
	betRepository := repository.NewBetRepository(userRepository, BETCOLLECTION)
	client := http.Client{}
	footballApi := api.NewFootBallApi(&client, Config.Api.BaseUrl, Config.Api.Token, Config.Api.Host)

	service.SetupUserService(userRepository)
	service.SetUpOddServivce(oddRepository, footballApi)
	service.SetupLeagueService(leagueRepository, footballApi)
	service.SetupMatchService(matchRepository, footballApi)
	service.SetupTeamService(teamRepository, footballApi)
	service.SetupBetService(betRepository)

	

}
