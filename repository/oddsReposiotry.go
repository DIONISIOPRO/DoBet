package repository

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var oddCollection = database.OpenCollection("odds")

type oddRepository struct{}

func NewOddRepository() OddRepository {
	repository := oddRepository{}
	return &repository
}

func (repo *oddRepository) UpSertOdd(odd models.Odds) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	_, err := oddCollection.InsertOne(ctx, odd)
	if err != nil {
		return err
	}
	return nil
}
func (repo *oddRepository) DeleteOdd(odd_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.D{{"league_id", odd_id}}

	_, err := oddCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *oddRepository) Odds(startIndex, perpage int64) ([]models.Odds, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var odds = []models.Odds{}
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex
	cursor, err := oddCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return odds, err
	}
	err = cursor.All(ctx, &odds)

	if err != nil {
		return odds, err
	}
	return odds, nil
}
