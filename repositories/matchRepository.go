package repositories

import (
	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
)

var matchCollection = database.OpenCollection("matches")

type matchRepository struct{

}

func NewMatchReposiotry() MatchRepository{
	return &matchRepository{}
}

func (repo *matchRepository) AddMatch(match models.Match) error {
	return nil
}

func (repo *matchRepository)DeleteMatch(match_id string) error {
	return nil
}

func (repo *matchRepository) UpDateMatch(match_id string, match models.Match) error {
	return nil
}

func (repo *matchRepository) Matches(startIndex, perpage int64) ([]models.Match, error ){
	return []models.Match{}, nil
}
