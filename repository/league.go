package repository

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type LeagueRepository interface {
	AddLeague(league models.League) error
	DeleteLeague(league_id string) error
	GetLeaguesByCountry(country string, startIndex, perpage int64) ([]models.League, error)
	Leagues(startIndex, perpage int64) ([]models.League, error)
}


type leagueRepository struct{
	Collection *mongo.Collection
}

func NewLeagueRepository(collection *mongo.Collection) LeagueRepository {
	return &leagueRepository{
		Collection: collection,
	}
}

func (repo *leagueRepository) AddLeague(league models.League) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	_, err := repo.Collection.InsertOne(ctx, league)
	if err != nil {
		return err
	}
	return nil
}

func (repo *leagueRepository) DeleteLeague(league_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.D{{"league_id", league_id}}

	_, err := repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *leagueRepository) Leagues(startIndex, perpage int64) ([]models.League, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *100)
	defer cancel()
	var allLeagues = []models.League{}
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex

	cursor, err := repo.Collection.Find(ctx, bson.D{{}}, opts)
	if err != nil{
		return allLeagues, err
	}
	err = cursor.All(ctx, &allLeagues)

	if err != nil{
		return allLeagues, err
	}
	return allLeagues, nil
}

func (repo *leagueRepository)GetLeaguesByCountry(country string, startIndex, perpage int64) ([]models.League, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *100)
	defer cancel()
	var allLeagues = []models.League{}
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex
	filter := bson.D{{"country", country}}

	cursor, err := repo.Collection.Find(ctx, filter, opts)
	if err != nil{
		return allLeagues, err
	}
	err = cursor.All(ctx, &allLeagues)

	if err != nil{
		return allLeagues, err
	}
	return allLeagues, nil
}

