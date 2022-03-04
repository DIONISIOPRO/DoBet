package repositories

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var betCollection = database.OpenCollection("bet")

func CreateBet(bet models.Bet) (bet_id string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	bet.ID = primitive.NewObjectID()
	bet_id = bet.ID.Hex()
	bet.Bet_id = bet_id

	_, insetErr := betCollection.InsertOne(ctx, bet)

	if insetErr != nil {
		return "", insetErr
	}
	return bet_id, nil
}

func BetByUser(user_id string) ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"bet_owner", user_id}}
	cursor, findErr := betCollection.Find(ctx, filter)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func BetByMatch(match_id string) ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"match_id", match_id}}
	cursor, findErr := betCollection.Find(ctx, filter)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func BetById(bet_id string) (models.Bet, error) {

	var bet models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"bet_id", bet_id}}
	cursor := betCollection.FindOne(ctx, filter)
	if err := cursor.Decode(bet); err != nil {
		return models.Bet{}, err
	}
	return bet, nil
}

func Bets() ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{}}
	cursor, findErr := betCollection.Find(ctx, filter)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func RunningBets() ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"isprocessed", false}}
	cursor, findErr := betCollection.Find(ctx, filter)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func TotalRunningBetsMoney() float32 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	type Data struct{
		total_count int64
	}

	var mydata Data

	filter := bson.D{{"$match", bson.D{{"isprocessed", false}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", "$amount"}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
	projectStage := bson.D{
		{
			"$project", bson.D{
				{"total_count", 1},
			}}}

	cursor, err := betCollection.Aggregate(ctx, mongo.Pipeline{
		filter, groupStage, projectStage,
	})

	if err != nil{
		return 0
	}
	errr := cursor.Decode(mydata)
	if errr != nil{
		return 0
	}

	return float32(mydata.total_count)
}

func UpdateBet(bet_id string, bet models.Bet) error{

	var updateOj primitive.D

	updateOj = append(updateOj, bson.E{"amount", bet.Amount} )
	updateOj = append(updateOj, bson.E{"potencial_win", bet.Amount} )
	updateOj = append(updateOj, bson.E{"isprocessed", bet.Amount} )
	updateOj = append(updateOj, bson.E{"isfinished", bet.Amount} )
	updateOj = append(updateOj, bson.E{"remain_matches", bet.Amount} )
	updateOj = append(updateOj, bson.E{"islose", bet.Amount} )

	filter := bson.M{"betid": bet_id}
	upsert := true
	options := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, updateErr := betCollection.UpdateOne(context.TODO(), filter,updateOj,&options )
	if updateErr != nil{
		return 	updateErr

	}
	return nil
}

func BetWatch(){
	
}




