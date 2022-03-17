package service

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

type OddService interface {
	UpSertOdd(odd models.Odds) error
	GetOddByMatchId(matchid string) (models.Odds, error )
	UpdateOdds(leagueId string) error
	DeleteOdd(odd_id string) error
	LunchUpdateOddsLoop()
}

type oddService struct {
	repository repository.OddRepository
	footballpi api.FootBallApi
	leagueservice LeagueService
}

func NewOddServivce(repository repository.OddRepository, footballpi api.FootBallApi, 	leagueservice LeagueService) OddService {
	return &oddService{
		repository: repository,
		footballpi: footballpi,
	}
}

func (service *oddService) UpSertOdd(odd models.Odds) error {
	validate := validator.New()
	err := validate.Struct(odd)
	if err != nil{
		return err
	}
	return service.repository.UpSertOdd(odd)

}
func (service *oddService) GetOddByMatchId(matchid string)( models.Odds, error ){
	if matchid == ""{
		return models.Odds{}, errors.New("Match Id Invalid")
	}
	return service.repository.GetOddByMatchId(matchid)
}
func (service *oddService) DeleteOdd(odd_id string) error {
	if odd_id == ""{
		return errors.New("Invalid odd id")
	}
	return service.repository.DeleteOdd(odd_id)
}

func (service *oddService) UpdateOdds(leagueId string) error {
	if leagueId == ""{
		return errors.New("Invalid league id")
	}
	oddDto, err := service.footballpi.GetOddsByLeagueId(leagueId)
	if err != nil {
		return err
	}
	odds := ConvertOddDtoToOddModelObjects(oddDto)

	requredGoroutines := len(odds)
	wg := &sync.WaitGroup{}
	wg.Add(requredGoroutines)
	for _, odd := range odds {
		func(odd models.Odds, wg *sync.WaitGroup) {
			defer wg.Done()
			service.repository.UpSertOdd(odd)
		}(odd, wg)
	}
	wg.Wait()
	return nil
}

func (service *oddService) LunchUpdateOddsLoop() {
	tiker := time.NewTicker(time.Hour * 24)
	leagues := []models.League{}

	for tker := range tiker.C {
		if len(LocalLeagues) == 0 {
			localLeagues, err := service.leagueservice.Leagues(0, 0)
			leagues = localLeagues
			if err != nil {
				return
			}

		} else {
			for _, league := range LocalLeagues {
				leagues = append(leagues, league)
			}

		}
		for _, league := range leagues {
			log := log.Default()
			log.Print(tker)
			go service.UpdateOdds(league.League_id)
		}

	}
}
