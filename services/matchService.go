package services

import (
	"log"

	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

var matchRepository repositories.MatchRepository

func NewMatchService(matchRepo repositories.MatchRepository) {
	matchRepository = matchRepo
}

func AddMatch(match models.Match) error {
	err := matchRepository.AddMatch(match)
	if err != nil {
		return err
	}
	provider := CreateBetProvider(match.Match_id)
	BetProviders[match.Match_id] = provider
	return nil
}

func DeleteMatch(match_id string) error {
	err := matchRepository.DeleteMatch(match_id)
	if err != nil {
		return err
	}
	delete(BetProviders, match_id)
	return nil
}

func UpDateMatch(match_id string, match models.Match) error {
	err := matchRepository.UpDateMatch(match_id, match)

	if err != nil {
		return err
	}
	provider := CreateBetProvider(match_id)
	BetProviders[match_id] = provider
	return nil
}

func Matches(startIndex, perpage int64) ([]models.Match, error) {
	return matchRepository.Matches(startIndex, perpage)
}

func MatchWatch() {
	UpdatedMtches, err := matchRepository.MatchWatch()
	if err != nil {
		log.Fatal(err)
	}
	for _, match := range UpdatedMtches {
		result := match.Result
		for _, bp := range BetProviders {
			if match.Match_id == bp.Match_id {
				bp.NotifyAll(result)
			}
		}
	}
}
