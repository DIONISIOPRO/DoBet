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
func (repo *paymentRepository) UpdateUser(userid string, user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"user_id": userid}
	_, err := repo.Collection.UpdateOne(ctx, filter, bson.M{"$set": user})
	if err != nil{
		return err
	}
	return nil
}

func (repo *paymentRepository) DeleteUser(user_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := User{}
	filter := bson.M{"user_id": user_id}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}
	if user.Balance > 0 {
		w := domain.WithDraw{
			User_id: user_id,
			Amount:  user.Balance,
		}
		repo.Withdraw(w)
	}
	_, err = repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentRepository) CheckMoney(amount float64, user_id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := User{}
	filter := bson.M{"user_id": user_id}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return false
	}
	return user.Balance >= amount
}

func (repo *paymentRepository) Deposit(deposit domain.Deposit) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"user_id": deposit.User_Id}
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
	filter := bson.M{"user_id": withdaw.User_id}
	canWithdraw := repo.CheckMoney(withdaw.Amount, withdaw.User_id)
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
