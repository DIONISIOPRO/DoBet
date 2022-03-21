package repository

import (
	"context"
	"fmt"
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BetRepository interface {
	CreateBet(bet models.Bet) (bet_id string, err error)
	UpdateBet(bet_id string, bet models.Bet) error
	BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error)
	BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error)
	BetById(bet_id string) (models.Bet, error)
	TotalBets() (int, error)
	TotalRunningBets() (int, error)
	Bets(startIndex, perpage int64) ([]models.Bet, error)
	RunningBets(startIndex, perpage int64) ([]models.Bet, error)
	TotalRunningBetsMoney() float64
	ProcessWin(amount float64, user_id string)
}

type betRepository struct {
	Collection         *mongo.Collection
	paymenteRepository PaymentRepository
}

func NewBetRepository(paymenteRepository PaymentRepository, Collection *mongo.Collection) BetRepository {
	return &betRepository{
		Collection:         Collection,
		paymenteRepository: paymenteRepository,
	}
}

func (repo *betRepository) CreateBet(bet models.Bet) (bet_id string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	bet.ID = primitive.NewObjectID()
	bet.Bet_id = bet.ID.Hex()
	bet_id = bet.Bet_id

	_, insetErr := repo.Collection.InsertOne(ctx, bet)
	if insetErr != nil {
		return "", insetErr
	}
	return bet_id, nil
}

func (repo *betRepository) BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
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

func (repo *betRepository) BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
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

func (repo *betRepository) BetById(bet_id string) (models.Bet, error) {
	var bet models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"bet_id": bet_id}
	cursor := repo.Collection.FindOne(ctx, filter)
	if err := cursor.Decode(bet); err != nil {
		return models.Bet{}, err
	}
	return models.Bet{}, nil
}

func (repo *betRepository) Bets(startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
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

func (repo *betRepository) TotalBets() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bets primitive.M
	filter := bson.D{primitive.E{Key: "$match", Value: bson.M{"isprocessed": true}}}
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

func (repo *betRepository) TotalRunningBets() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bets primitive.M
	filter := bson.D{primitive.E{Key: "$match", Value: bson.M{"isprocessed": false}}}
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

func (repo *betRepository) RunningBets(startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"isprocessed": false}
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

func (repo *betRepository) TotalRunningBetsMoney() float64 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var bets []primitive.M
	filter := bson.D{primitive.E{Key: "$match", Value: bson.M{"isprocessed": false}}}
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

func (repo *betRepository) UpdateBet(bet_id string, bet models.Bet) error {
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

func (repo *betRepository) ProcessWin(amount float64, user_id string) {
	repo.paymenteRepository.Deposit(amount, user_id)
}
