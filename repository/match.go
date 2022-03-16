package repository

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type MatchRepository interface {
	DeleteOldMatch() error
	UpDateMatch(match_id string, match models.Match) error
	MatchesByLeagueId(leagueId string, startIndex, perpage int64) ([]models.Match, error)
	MatchesByLeagueIdDay(leagueid string, day, startIndex, perpage int64) ([]models.Match, error)
	Matches(startIndex, perpage int64) ([]models.Match, error)
	MatchWatch(f func(models.Match))
}
type matchRepository struct {
	Collection *mongo.Collection
}

func NewMatchReposiotry(collectionName string) MatchRepository {
	matchCollection := database.OpenCollection(collectionName)
	return &matchRepository{
		Collection: matchCollection,
	}
}

func (repo *matchRepository) DeleteOldMatch() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	now := time.Now().Unix()
	monthBefore := time.Now().Add(-time.Second * 60 * 60 * 24 * 30)
	diference := now - monthBefore.Unix()
	filter := bson.D{{"time", bson.D{{"$lt", diference}}}}
	_, err := repo.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepository) UpDateMatch(match_id string, match models.Match) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}
	filter := bson.D{{"match_id", match_id}}
	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{{"$set", match}}, &opts)
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepository) Matches(startIndex, perpage int64) ([]models.Match, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var allMatches []models.Match
	opts := options.Find()
	opts.SetLimit(perpage)
	opts.SetSkip(startIndex)

	cursor, err := repo.Collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return allMatches, err
	}
	err = cursor.All(ctx, &allMatches)
	if err != nil {
		return allMatches, err
	}
	return allMatches, nil
}

func (repo *matchRepository) MatchesByLeagueId(leagueId string, startIndex, perpage int64) ([]models.Match, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var allMatches []models.Match
	filter := bson.D{{"league", leagueId}}
	opts := options.Find()
	opts.SetLimit(perpage)
	opts.SetSkip(startIndex)

	cursor, err := repo.Collection.Find(ctx, filter, opts)
	if err != nil {
		return allMatches, err
	}
	err = cursor.All(ctx, &allMatches)
	if err != nil {
		return allMatches, err
	}
	return allMatches, nil
}

func (repo *matchRepository) MatchesByLeagueIdDay(leagueid string, day, startIndex, perpage int64) ([]models.Match, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var allMatches []models.Match
	oneDay := 86400000000000
	now := time.Now().Unix()
	days := int64(oneDay) * day
	remainDays := now + days
	leagueFilter := bson.E{"league", leagueid}
	lessDaysFilter := bson.E{"day", bson.D{{"$lt", remainDays+(remainDays/2)}}}
	greaterDaysFilter := bson.E{"day", bson.D{{"$gt", remainDays-(remainDays/2)}}}
	filter := bson.D{leagueFilter,lessDaysFilter, greaterDaysFilter}
	opts := options.Find()
	opts.SetLimit(perpage)
	opts.SetSkip(startIndex)

	cursor, err := repo.Collection.Find(ctx, filter, opts)
	if err != nil {
		return allMatches, err
	}
	err = cursor.All(ctx, &allMatches)
	if err != nil {
		return allMatches, err
	}
	return allMatches, nil}

func (repo *matchRepository) MatchWatch(f func(models.Match)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var allMatches []models.Match
	matchStage := bson.D{{"$match", bson.D{{"operationType", "update"}}}}
	projectStage := bson.D{{"$project", bson.D{{"fullDocument", 1}}}}
	opts := options.ChangeStream().SetMaxAwaitTime(2 * time.Second)
	stream, err := repo.Collection.Watch(ctx, mongo.Pipeline{
		matchStage, projectStage,
	}, opts)

	if err != nil {
		panic(err)
	}
	err = stream.Decode(&allMatches)

	if err != nil {
		panic(err)
	}
	for _, m := range allMatches {
		go f(m)
	}
}
