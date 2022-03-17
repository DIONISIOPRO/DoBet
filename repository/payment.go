package repository

import (
	"context"
	"errors"
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository interface {
	Deposit(amount float64, userid string) error
	Withdraw(amount float64, userid string) error
}
type paymenteRepository struct {
	Collection *mongo.Collection
}

func NewPaymentReposiotry(Usercollection *mongo.Collection) PaymentRepository {
	return &paymenteRepository{
		Collection: Usercollection,
	}
}

func (repo *paymenteRepository) Deposit(amount float64, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{{"user_id", userid}}
	var user = models.User{}

	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}
	updatedBalance := user.Account_balance + amount
	object := bson.E{"account_balance", updatedBalance}


	_, err = repo.Collection.UpdateOne(ctx, filter, bson.D{{"$set", object}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymenteRepository) Withdraw(amount float64, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{{"user_id", userid}}
	var user = models.User{}

	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}
	if user.Account_balance < amount{
		return errors.New("balance low than amount")
	}
	currentBalance := user.Account_balance - amount
	updateObj := bson.E{"account_balance", currentBalance}
	_, err = repo.Collection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}})
	if err != nil {
		return err
	}
	return nil
}
