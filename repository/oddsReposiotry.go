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


type oddRepository struct{
	Collection *mongo.Collection
}

func NewOddRepository(collectioName string) OddRepository {
	oddCollection := database.OpenCollection(collectioName)
	repository := oddRepository{
		Collection: oddCollection,
	}
	return &repository
}

func (repo *oddRepository) UpSertOdd(odd models.Odds) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{{"match_id", odd.Match_id}}
	Upsert := true
	opts := &options.UpdateOptions{
		Upsert: &Upsert,
	}

	_, err := repo.Collection.UpdateOne(ctx, filter, odd, opts)
	if err != nil {
		return err
	}
	return nil
}
func (repo *oddRepository) DeleteOdd(odd_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.D{{"odd_id", odd_id}}

	_, err := repo.Collection.DeleteOne(ctx, filter)
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
	cursor, err := repo.Collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return odds, err
	}
	err = cursor.All(ctx, &odds)

	if err != nil {
		return odds, err
	}
	return odds, nil
}

func (repo *oddRepository) GetOddByMatchId(match_id string) (models.Odds) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	filter := bson.D{{"match_id", match_id}}
	odd := models.Odds{}
	repo.Collection.FindOne(ctx,filter).Decode(&odd)
	return odd
}
