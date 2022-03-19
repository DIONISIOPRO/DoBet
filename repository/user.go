package repository

import (
	"context"
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetUserById(userId string) (models.User, error)
	GetUserByPhone(phone string) (models.User, error)
	Users(startIndex, perpage int64) ([]models.User, error)
}

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		Collection: collection,
	}
}


func (repo *userRepository) Users(startIndex, perpage int64) ([]models.User, error) {
	allusers := []models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex

	cursor, err := repo.Collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return []models.User{}, err
	}
	err = cursor.All(ctx, &allusers)
	if err != nil {
		return []models.User{}, err
	}
	return allusers, nil
}

func (repo *userRepository) GetUserById(userId string) (models.User, error) {
	user := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"user_id": userId}

	cursor, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		return user, err
	}
	err = cursor.All(ctx, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) GetUserByPhone(phone string) (models.User, error) {
	user := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"phone_number": phone}

	cursor, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		return user, err
	}
	err = cursor.All(ctx, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
