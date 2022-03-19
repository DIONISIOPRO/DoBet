package repository

import (
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var MatchCollection = []models.Match{}

type MatchRepository interface {
	DeleteOldMatch()
	UpDateMatch(match_id string, match models.Match)
	MatchesByLeagueIdDay(leagueid string, day, startIndex, perpage int64) ([]models.Match, error)
}
type matchRepository struct {
	Collection *mongo.Collection
}

func NewMatchReposiotry(collection *mongo.Collection) MatchRepository {
	return &matchRepository{
		Collection: collection,
	}
}

func (repo *matchRepository) DeleteOldMatch() {
	now := time.Now().Unix()
	monthBefore := time.Now().Add(-time.Second * 60 * 60 * 24 * 30)
	diference := now - monthBefore.Unix()

	for _, m := range MatchCollection {
		if m.Time < diference {
			MatchCollection = append(MatchCollection, m)
		}
	}

}

func (repo *matchRepository) UpDateMatch(match_id string, match models.Match) {
	MatchCollection = append(MatchCollection, match)
}

func (repo *matchRepository) MatchesByLeagueIdDay(leagueid string, day, startIndex, perpage int64) ([]models.Match, error) {
	oneDay := 86400000000000
	now := time.Now().Unix()
	days := int64(oneDay) * day
	remainDays := now + days
	lessday := remainDays + (remainDays / 2)
	greaterday := remainDays - (remainDays / 2)

	matches := []models.Match{}
	for _, m := range MatchCollection {
		if m.LeagueId == leagueid && m.Time > greaterday && m.Time < lessday {
			matches = append(matches, m)
		}
	}
	if len(matches) > int(startIndex) {
		if len(matches) >= int(startIndex + perpage) {
			return matches[startIndex : startIndex + perpage], nil
		} else {
			return matches[startIndex : len(matches) - 1], nil

		}
	}

	return []models.Match{}, nil
}
