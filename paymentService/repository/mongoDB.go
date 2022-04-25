package repository

import (
	"context"
	"errors"
	"time"
	"github.com/dionisiopro/dobet_payment/domain"
)

type User struct {
	User_id      string  `bson:"user_id"`
	Phone_number string  `bson:"phone_number"`
	Balance      float64 `bson:"balance"`
}
type MongoDriverUserCollection interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}
type paymentRepository struct {
	Collection MongoDriverUserCollection
}

func NewPaymentMongodbReposiotry(Usercollection MongoDriverUserCollection) *paymentRepository {
	return &paymentRepository{
		Collection: Usercollection,
	}
}
func (repo *paymentRepository) CreateUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := repo.Collection.InsertOne(ctx, user)
	if err != nil {
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
		repo.Withdraw(user.Balance, user_id)
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
func (repo *paymentRepository) Deposit(amount float64, user_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"user_id": user_id}
	updateObj := bson.D{{Key: "$inc", Value: bson.D{{Key: "balance", Value: amount}}}}
	_, err := repo.Collection.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentRepository) Withdraw(amount float64, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"user_id": userid}
	canWithdraw  := repo.CheckMoney(amount, userid)
	if !canWithdraw{
		return errors.New(domain.NotEnoughMoney)
	}
	updateObj := bson.D{{Key: "$inc", Value: bson.D{{Key: "balance", Value: -amount}}}}
	_, err := repo.Collection.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return err
	}
	return nil
}
