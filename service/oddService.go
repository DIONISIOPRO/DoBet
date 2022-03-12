package service

import (
	"log"
	"sync"
	"time"

	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/models"
	"gitthub.com/dionisiopro/dobet/repository"
)

var OddService = &oddService{}

type oddService struct {
	repository repository.OddRepository
	footballpi api.FootBallApi
}

func SetUpOddServivce(repository repository.OddRepository,footballpi api.FootBallApi ){
	OddService.footballpi = footballpi
	OddService.repository = repository
	lunchUpdateOddsLoop()
}

func (service *oddService) UpSertOdd(odd models.Odds) error {
	return service.repository.UpSertOdd(odd)

}
func (service *oddService) GetOddByMatchId(matchid string) (models.Odds){
	return service.repository.GetOddByMatchId(matchid)
}
func (service *oddService) DeleteOdd(odd_id string) error {
	return service.repository.DeleteOdd(odd_id)
}

func (service *oddService) UpdateOdds(leagueId string) error {
	oddDto, err := OddService.footballpi.GetOddsByLeagueId(leagueId)
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
			OddService.repository.UpSertOdd(odd)
		}(odd, wg)
	}
	wg.Wait()
	return nil
}

func lunchUpdateOddsLoop() {
	tiker := time.NewTicker(time.Hour * 24)
	leagues := []models.League{}

	for tker := range tiker.C {
		if len(LocalLeagues) == 0 {
			localLeagues, err := LeagueService.Leagues(0, 0)
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
			go OddService.UpdateOdds(league.League_id)
		}

	}
}
