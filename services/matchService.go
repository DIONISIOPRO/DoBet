package services

import (
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

func AddMatch(match models.Match) error {
	 err := repositories.AddMatch(match)
	 if err != nil{
		 return err
	 }
	 provider := CreateBetProvider(match.Match_id)
	 BetProviders[match.Match_id] = provider
	 return nil
}

func DeleteMatch(match_id string) error {
	err :=  repositories.DeleteMatch(match_id)
	if err != nil{
		return err
	}
	delete(BetProviders, match_id)
	return nil
}

func UpDateMatch(match_id string, match models.Match) error {
	err := repositories.UpDateMatch(match_id, match)

	if err != nil{
		return err
	}
	provider := CreateBetProvider(match_id)
	BetProviders[match_id] = provider
	return nil
}

func Matches() []models.Match {
	return repositories.Matches()
}