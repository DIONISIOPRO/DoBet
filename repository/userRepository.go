package repository

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

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collectionName string) UserRepository {
	userColletion := database.OpenCollection(collectionName)
	return &userRepository{
		Collection: userColletion,
	}
}


func (repo *userRepository) Deposit(amount float64, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{{"user_id", userid}}
	var updateObj primitive.M

	cursor, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	err = cursor.Decode(updateObj)
	if err != nil {
		return err
	}

	updateObj["amount"] = updateObj["amount"].(float64) + amount

	_, err = repo.Collection.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Withdraw(amount float64, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{{"user_id", userid}}
	var updateObj primitive.M

	cursor, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	err = cursor.Decode(updateObj)
	if err != nil {
		return err
	}
	currentBalace := updateObj["amount"].(float64)
	if currentBalace >= amount {
		updateObj["amount"] = updateObj["amount"].(float64) - amount
	}
	_, err = repo.Collection.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) SignUp(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	_, err := repo.Collection.InsertOne(ctx,user)
	if err != nil {
		return err
	}
	return nil
}
func (repo *userRepository) Login(user models.User) (models.User, error){
return models.User{}, nil
}


func (repo *userRepository) Users(startIndex, perpage int64) ([]models.User, error) {
	allusers := []models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

	opts := options.Find()
	opts.Limit =  &perpage
	opts.Skip = &startIndex

	cursor, err := repo.Collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return []models.User{},err
	}
	err = cursor.All(ctx, &allusers)
	if err != nil {
		return []models.User{},err
	}
	return allusers, nil
}
