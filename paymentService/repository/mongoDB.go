package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dionisiopro/dobet_payment/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	User_id      string  `bson:"user_id"`
	Phone_number string  `bson:"phone_number"`
	Balance      float64 `bson:"balance"`
}

type paymentRepository struct {
	Collection *mongo.Collection
}

func NewPaymentMongodbReposiotry(collection *mongo.Collection) *paymentRepository {
	return &paymentRepository{
		Collection: collection,
	}
}

func (repo *paymentRepository) CreateUser(user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := repo.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentRepository) DeleteUser(phone_number string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := User{}
	filter := bson.M{"phone_number": phone_number}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}
	if user.Balance > 0 {
		w := domain.WithDraw{
			Phone_number: phone_number,
			Amount:       user.Balance,
		}
		repo.Withdraw(w)
	}
	_, err = repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentRepository) CheckMoney(amount float64, phone_number string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := User{}
	filter := bson.M{"phone_number": phone_number}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return false
	}
	return user.Balance >= amount
}
func (repo *paymentRepository) Deposit(deposit domain.Deposit) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"phone_number": deposit.Phone_number}
	updateObj := bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "balance", Value: deposit.Amount}}}}
	_, err := repo.Collection.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentRepository) Withdraw(withdaw domain.WithDraw) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"phone_number": withdaw.Phone_number}
	canWithdraw := repo.CheckMoney(withdaw.Amount, withdaw.Phone_number)
	if !canWithdraw {
		return errors.New(domain.NotEnoughMoney)
	}
	updateObj := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "balance", Value: -withdaw.Amount}}}}
	_, err := repo.Collection.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return err
	}
	return nil
}
