package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/dionisiopro/dobet-bet/domain/bet"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BetReposiotry struct {
	Collection *mongo.Collection
}

func NewBetReposiotry(Collection *mongo.Collection) *BetReposiotry {
	return &BetReposiotry{
		Collection: Collection,
	}
}

func (repo *BetReposiotry) CreateBet(bet bet.BetBase) (bet_id string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	id := primitive.NewObjectID()
	bet.Bet_id = id.Hex()
	bet_id = bet.Bet_id

	_, insetErr := repo.Collection.InsertOne(ctx, bet)
	if insetErr != nil {
		return "", insetErr
	}
	return bet_id, nil
}

func (repo *BetReposiotry) ConfirmBet(bet_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	filter := bson.D{{Key: "bet_id", Value: bet_id}}
	updateObj := bson.E{Key: "status", Value: bet.Confirmed}
	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: updateObj}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *BetReposiotry) CancelBet(bet_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	filter := bson.D{{Key: "bet_id", Value: bet_id}}
	updateObj := bson.E{Key: "status", Value: bet.Canceled}
	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: updateObj}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *BetReposiotry) ActiveBet(bet_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	filter := bson.D{{Key: "bet_id", Value: bet_id}}
	updateObj := bson.E{Key: "status", Value: bet.Active}
	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: updateObj}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *BetReposiotry) AllRunningBetsByMatch(match_id string) ([]bet.BetBase, error) {
	var allbets []bet.BetBase
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"match_id": match_id, "is_finished": false}

	cursor, findErr := repo.Collection.Find(ctx, filter)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func (repo *BetReposiotry) BetByUser(user_id string, startIndex, perpage int64) ([]bet.BetBase, error) {
	var allbets []bet.BetBase
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)
	defer cancel()
	filter := bson.M{"bet_owner": user_id}
	cursor, findErr := repo.Collection.Find(ctx, filter, opts)
	if findErr != nil {
		return allbets, findErr
	}
	err := cursor.All(ctx, &allbets)
	if err != nil {
		fmt.Print("error here")
	}
	return allbets, nil
}

func (repo *BetReposiotry) BetByMatch(match_id string, startIndex, perpage int64) ([]bet.BetBase, error) {
	var allbets []bet.BetBase
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"match_id": match_id}
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)
	cursor, findErr := repo.Collection.Find(ctx, filter, opts)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func (repo *BetReposiotry) BetById(bet_id string) (bet.BetBase, error) {
	var _bet bet.BetBase
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"bet_id": bet_id}
	cursor := repo.Collection.FindOne(ctx, filter)
	if err := cursor.Decode(_bet); err != nil {
		return bet.BetBase{}, err
	}
	return bet.BetBase{}, nil
}

func (repo *BetReposiotry) Bets(startIndex, perpage int64) ([]bet.BetBase, error) {
	var allbets []bet.BetBase
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{}
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)

	cursor, findErr := repo.Collection.Find(ctx, filter, opts)
	if findErr != nil {
		return allbets, findErr
	}
	if err := cursor.All(ctx, &allbets); err != nil {
		return allbets, err

	}
	return allbets, nil
}

func (repo *BetReposiotry) TotalBets() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bets primitive.M
	filter := bson.D{primitive.E{Key: "$match", Value: bson.M{"is_finished": true}}}
	group := bson.D{primitive.E{Key: "$group", Value: bson.M{"_id": ""}}, primitive.E{Key: "totalbets", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}
	project := bson.D{primitive.E{Key: "$project", Value: bson.D{primitive.E{Key: "$totalbets", Value: 1}}}}
	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{filter, group, project})

	if err != nil {
		return -1, err
	}
	err = cursor.All(ctx, &bets)

	if err != nil {
		return -1, err
	}
	total := bets["totalbets"]
	return total.(int), nil
}

func (repo *BetReposiotry) TotalRunningBets() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bets primitive.M
	filter := bson.D{primitive.E{Key: "$match", Value: bson.M{"is_finished": false}}}
	group := bson.D{primitive.E{Key: "&group", Value: bson.D{primitive.E{Key: "_id", Value: ""}, primitive.E{Key: "totalbets", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}
	project := bson.D{primitive.E{Key: "$project", Value: bson.D{primitive.E{Key: "$totalbets", Value: 1}}}}

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{filter, group, project})

	if err != nil {
		return -1, err
	}
	err = cursor.All(ctx, &bets)

	if err != nil {
		return -1, err
	}
	total := bets["totalbets"]
	return total.(int), nil
}

func (repo *BetReposiotry) RunningBets(startIndex, perpage int64) ([]bet.BetBase, error) {
	var allbets []bet.BetBase
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"is_finished": false}
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)
	cursor, findErr := repo.Collection.Find(ctx, filter, opts)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func (repo *BetReposiotry) TotalRunningBetsMoney() float64 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var bets []primitive.M
	filter := bson.D{primitive.E{Key: "$match", Value: bson.M{"is_finished": false}}}
	group := bson.D{primitive.E{Key: "&group", Value: bson.D{primitive.E{Key: "_id", Value: ""}, primitive.E{Key: "totalMoney", Value: primitive.E{Key: "$sum", Value: "$amount"}}}}}
	projectStage := bson.D{primitive.E{Key: "$project", Value: bson.M{"$totalMoney": 1}}}
	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{filter, group, projectStage})
	if err != nil {
		return -1
	}
	errr := cursor.All(ctx, &bets)
	if errr != nil {
		return -1
	}
	var money = bets[0]["totalMoney"]
	return money.(float64)
}

func (repo *BetReposiotry) UpdateBet(bet_id string, bet bet.BetBase) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"bet_id": bet_id}
	upsert := true
	options := options.UpdateOptions{
		Upsert: &upsert,
	}
	var updateObj primitive.D
	betDoc, err := bson.Marshal(bet)
	if err != nil {
		return err
	}
	if err = bson.Unmarshal(betDoc, &updateObj); err != nil {
		return err
	}
	_, updateErr := repo.Collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: updateObj}}, &options)
	if updateErr != nil {
		return updateErr
	}
	return nil
}
