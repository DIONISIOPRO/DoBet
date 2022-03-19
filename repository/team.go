package repository

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamRepository interface {
	Upsert(team models.Team) error
	DeleteTeam(team_id string) error
	TeamsByCountry(country string, startIndex, perpage int64) ([]models.Team, error)
	Teams(startIndex, perpage int64) ([]models.Team, error)
}
type teamRepository struct {
	Collection *mongo.Collection
}

func NewTeamRepository(collection *mongo.Collection) TeamRepository {
	return &teamRepository{
		Collection: collection,
	}
}

func (repo *teamRepository) Upsert(team models.Team) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	filter := bson.M{"team_id":team.Team_id}
	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{{"$set", team}}, &opts)
	if err != nil {
		return err
	}
	return nil
}

func (repo *teamRepository) TeamsByCountry(country string, startIndex, perpage int64) ([]models.Team, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var allTeams []models.Team
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex
	filter := bson.M{"country": country}

	cursor, err := repo.Collection.Find(ctx, filter, opts)
	if err != nil {
		return allTeams, err
	}
	err = cursor.All(ctx, &allTeams)

	if err != nil {
		return allTeams, err
	}
	return allTeams, nil
}

func (repo *teamRepository) DeleteTeam(team_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	filter := bson.M{"team_id": team_id}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *teamRepository) Teams(startIndex, perpage int64) ([]models.Team, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var allTeams []models.Team
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex

	cursor, err := repo.Collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return allTeams, err
	}
	err = cursor.All(ctx, &allTeams)

	if err != nil {
		return allTeams, err
	}
	return allTeams, nil
}
