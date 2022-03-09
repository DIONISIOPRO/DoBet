package repository

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var teamCollection = database.OpenCollection("teams")

type teamRepository struct{

}

func NewTeamRepository() TeamRepository{
	return &teamRepository{}
}


func (repo *teamRepository) AddTeam(team models.Team) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *100)
	defer cancel()

	_, err := teamCollection.InsertOne(ctx, team)
	if err != nil{
		return err
	}
	return nil
}

func (repo *teamRepository)  DeleteTeam(team_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *100)
	defer cancel()

	filter := bson.M{"team_id":team_id}
	_, err := teamCollection.DeleteOne(ctx, filter)
	if err != nil{
		return err
	}
	return nil
}


func(repo *teamRepository)  Teams(startIndex, perpage int64) ([]models.Team, error ){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *100)
	defer cancel()
	var allTeams []models.Team
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex

	cursor, err := teamCollection.Find(ctx, bson.D{{}}, opts)
	if err != nil{
		return allTeams, err
	}
	err = cursor.All(ctx, &allTeams)

	if err != nil{
		return allTeams, err
	}
	return allTeams, nil
}
