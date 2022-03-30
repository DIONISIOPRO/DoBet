package repository

import (
	"context"
	"time"

	"github/namuethopro/dobet-user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetUserById(userId string) (domain.User, error)
	GetUserByPhone(phone string) (domain.User, error)
	Users(startIndex, perpage int64) ([]domain.User, error)
	DeleteUser(userid string) error
	UpdateUser(userid string) error
}

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		Collection: collection,
	}
}

func (repo *userRepository) Users(startIndex, perpage int64) ([]domain.User, error) {
	users := []domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex
	cursor, err := repo.Collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return []domain.User{}, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (repo *userRepository) GetUserById(userId string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	
	err := repo.Collection.FindOne(ctx, idfilter(userId)).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) UpdateUser(userId string) error {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	
	_, err := repo.Collection.UpdateOne(ctx, idfilter(userId), bson.D{{Key: "$set", Value: user}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUserByPhone(phone string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"phone_number": phone}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) DeleteUser(userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := repo.Collection.DeleteOne(ctx, idfilter(userid))
	if err != nil {
		return err
	}
	return nil
}

func idfilter(id string) (filter bson.M){
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		filter = bson.M{"user_id": id}
		return
	}
	filter = bson.M{"_id": _id}
	return
}
