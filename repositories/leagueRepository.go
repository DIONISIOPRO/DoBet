package repositories

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var leagueCollection = database.OpenCollection("leagues")

type leagueRepository struct{}

func NewLeagueRepository() LeagueRepository {
	return &leagueRepository{}
}

func (service *leagueRepository) AddLeague(league models.League) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	_, err := leagueCollection.InsertOne(ctx, league)
	if err != nil {
		return err
	}
	return nil
}

func (service *leagueRepository) DeleteLeague(league_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.D{{"league_id", league_id}}

	_, err := leagueCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (service *leagueRepository) Leagues(startIndex, perpage int64) ([]models.League, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *100)
	defer cancel()
	var allLeagues []models.League
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex

	cursor, err := leagueCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil{
		return allLeagues, err
	}
	err = cursor.All(ctx, &allLeagues)

	if err != nil{
		return allLeagues, err
	}
	return allLeagues, nil
}
