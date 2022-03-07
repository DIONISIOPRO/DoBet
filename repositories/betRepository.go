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

var betCollection = database.OpenCollection("bets")

type betRepository struct {
	userRepository UserRepository
}

func NewBetRepository(userRepository UserRepository) BetRepository {
	return &betRepository{
		userRepository: userRepository,
	}
}

func (repo *betRepository) CreateBet(bet models.Bet) (bet_id string, err error) {
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

func (repo *betRepository) BetByUser(user_id string, startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)
	defer cancel()
	filter := bson.D{{"bet_owner", user_id}}
	cursor, findErr := betCollection.Find(ctx, filter, opts)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func (repo *betRepository) BetByMatch(match_id string, startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"match_id", match_id}}
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)
	cursor, findErr := betCollection.Find(ctx, filter, opts)
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
	filter := bson.D{{"bet_id", bet_id}}
	cursor := betCollection.FindOne(ctx, filter)
	if err := cursor.Decode(bet); err != nil {
		return models.Bet{}, err
	}
	return bet, nil
}

func (repo *betRepository) Bets(startIndex, perpage int64) ([]models.Bet, error) {
	var allbets []models.Bet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)

	cursor, findErr := betCollection.Find(ctx, filter, opts)
	if findErr != nil {
		return allbets, findErr
	}
	cursor.All(ctx, allbets)
	return allbets, nil
}

func (repo *betRepository) TotalBets() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bets primitive.M
	filter := bson.D{{"$match", bson.D{{"isprocessed", true}}}}
	count := bson.D{{"totalbets", bson.D{{"$sum", 1}}}}
	project := bson.D{{"$project", bson.D{{"$totalbets", 1}}}}

	cursor, err := betCollection.Aggregate(ctx, mongo.Pipeline{filter, count, project})

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
	filter := bson.D{{"$match", bson.D{{"isprocessed", false}}}}
	count := bson.D{{"totalbets", bson.D{{"$sum", 1}}}}
	project := bson.D{{"$project", bson.D{{"$totalbets", 1}}}}

	cursor, err := betCollection.Aggregate(ctx, mongo.Pipeline{filter, count, project})

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
	filter := bson.D{{"isprocessed", false}}
	opts := options.Find()
	opts.SetSkip(startIndex)
	opts.SetLimit(perpage)
	cursor, findErr := betCollection.Find(ctx, filter, opts)
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
	filter := bson.D{{"$match", bson.D{{"isprocessed", false}}}}
	sumStage := bson.D{{"totalMoney", bson.M{"$sum": "$amount"}}}
	projectStage := bson.D{{"$project", bson.M{"$totalMoney": 1}}}
	cursor, err := betCollection.Aggregate(ctx, mongo.Pipeline{filter, sumStage, projectStage})
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
	_, updateErr := betCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &options)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (repo *betRepository) ProcessWin(amount float64, user_id string) {
	repo.userRepository.Deposit(amount, user_id)
}
