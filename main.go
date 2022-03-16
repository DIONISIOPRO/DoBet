package main

import (
	"encoding/json"
	"os"

	"gitthub.com/dionisiopro/dobet/config"
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

}
