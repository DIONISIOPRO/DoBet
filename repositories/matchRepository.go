package repositories

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

func (repo *matchRepository) AddMatch(match models.Match) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	_, err := matchCollection.InsertOne(ctx, match)
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepository) DeleteMatch(match_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.D{{"match_id", match_id}}
	_, err := matchCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepository) UpDateMatch(match_id string, match models.Match) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.D{{"match_id", match_id}}
	_, err := matchCollection.UpdateOne(ctx, filter, match)
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

func (repo *matchRepository) MatchWatch() ([]models.Match, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var allMatches []models.Match
	matchStage := bson.D{{"$match", bson.D{{"operationType", "update"}}}}
	opts := options.ChangeStream().SetMaxAwaitTime(2 * time.Second)
	stream, err := matchCollection.Watch(ctx, mongo.Pipeline{
		matchStage,
	}, opts)

	if err != nil {
		return allMatches, err
	}
	err = stream.Decode(&allMatches)

	if err != nil {
		return allMatches, err
	}

	return allMatches, nil
}
