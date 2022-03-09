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

var matchCollection = database.OpenCollection("matches")

type matchRepository struct {
}

func NewMatchReposiotry() MatchRepository {
	return &matchRepository{}
}

func (repo *matchRepository) DeleteOldMatch() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	now := time.Now().Unix()
	monthBefore := time.Now().Add(-time.Second * 60 * 60 * 24 * 30)
	diference := now - monthBefore.Unix()
	filter := bson.D{{"time", bson.D{{"$lt", diference}}}}
	_, err := matchCollection.DeleteMany(ctx, filter)
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
	_, err := matchCollection.UpdateOne(ctx, filter, match, &opts)
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

	cursor, err := matchCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return allMatches, err
	}
	err = cursor.All(ctx, &allMatches)
	if err != nil {
		return allMatches, err
	}
	return allMatches, nil
}

func (repo *matchRepository) MatchWatch(f func(models.Match)){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var allMatches []models.Match
	matchStage := bson.D{{"$match", bson.D{{"operationType", "update"}}}}
	projectStage := bson.D{{"$project", bson.D{{"fullDocument", 1}}}}
	opts := options.ChangeStream().SetMaxAwaitTime(2 * time.Second)
	stream, err := matchCollection.Watch(ctx, mongo.Pipeline{
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
