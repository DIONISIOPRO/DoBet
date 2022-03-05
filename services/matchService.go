package services

import (
	"log"

	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repositories"
)

type matchService struct{
	repo repositories.MatchRepository
	betService BetService
}

func NewMatchService(matchRepo repositories.MatchRepository, betService BetService) MatchService{
 return &matchService{
	 repo:  matchRepo,
	 betService: betService ,
 }

}

func (service *matchService) AddMatch(match models.Match) error {
	 err := service.repo.AddMatch(match)
	 if err != nil{
		 return err
	 }
	 provider := service.betService.CreateBetProvider(match.Match_id)
	 BetProviders[match.Match_id] = provider
	 return nil
}

func (service *matchService)DeleteMatch(match_id string) error {
	err :=  service.repo.DeleteMatch(match_id)
	if err != nil{
		return err
	}
	delete(BetProviders, match_id)
	return nil
}

func (service *matchService)UpDateMatch(match_id string, match models.Match) error {
	err := service.repo.UpDateMatch(match_id, match)

	if err != nil{
		return err
	}
	provider := service.betService.CreateBetProvider(match_id)
	BetProviders[match_id] = provider
	return nil
}

func(service *matchService) Matches(startIndex, perpage int64) ([]models.Match, error) {
	return service.repo.Matches(startIndex, perpage)
}

func(service *matchService) MatchWatch(){
	UpdatedMtches, err:= service.repo.MatchWatch()
	if err != nil{
		log.Fatal(err)
	}
	for _, match := range UpdatedMtches {
		result := match.Result
		for _, bp := range BetProviders {
			if match.Match_id == bp.Match_id{
				bp.NotifyAll(result)
			}
		}
	}
}