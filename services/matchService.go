package services

import (
	"log"

	"gitthub.com/dionisiopro/dobet/helpers"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)


var MatchService matchService

type matchService struct {
	repo repositories.MatchRepository
}

func SetupMatchService(matchRepo repositories.MatchRepository) *matchService {
	MatchService.repo = matchRepo
	return  &MatchService
}

func (service *matchService)AddMatch(match models.Match) error {
	err := service.repo.AddMatch(match)
	if err != nil {
		return err
	}
	provider := helpers.CreateBetProvider(match.Match_id)
	BetProviders[match.Match_id] = provider
	return nil
}

func (service *matchService)DeleteMatch(match_id string) error {
	err := service.repo.DeleteMatch(match_id)
	if err != nil {
		return err
	}
	delete(BetProviders, match_id)
	return nil
}

func (service *matchService)UpDateMatch(match_id string, match models.Match) error {
	err := service.repo.UpDateMatch(match_id, match)

	if err != nil {
		return err
	}
	provider := helpers.CreateBetProvider(match_id)
	BetProviders[match_id] = provider
	return nil
}

func (service *matchService)Matches(startIndex, perpage int64) ([]models.Match, error) {
	return service.repo.Matches(startIndex, perpage)
}

func (service *matchService) MatchWatch() {
	UpdatedMtches, err := service.repo.MatchWatch()
	if err != nil {
		log.Fatal(err)
	}
	for _, match := range UpdatedMtches {
		result := match.Result
		for _, bp := range BetProviders {
			if match.Match_id == bp.Match_id {
				bp.NotifyAll(result, BetService.ProcessBet)
			}
		}
	}
}
