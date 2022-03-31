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
	SetIndexes()
	GetUserById(userId string) (domain.User, error)
	GetUserByPhone(phone string) (domain.User, error)
	Users(startIndex, perpage int64) ([]domain.User, error)
	DeleteUser(userid string) error
	UpdateUser(userid string, user domain.User) error
}

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	repo :=  &userRepository{
		Collection: collection,
	}
	repo.SetIndexes()
	return repo
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

func (repo *userRepository) UpdateUser(userId string, user domain.User) error {
	updateObj := bson.M{}
	if user.User_id != "" {
		updateObj["user_id"] = user.User_id
	}
	if user.First_name != "" {
		updateObj["first_name"] = user.First_name
	}
	if user.Last_name != "" {
		updateObj["last_name"] = user.First_name
	}
	if user.Phone_number != "" {
		updateObj["phone_number"] = user.Phone_number

	}
	if user.Hashed_password != "" {
		updateObj["hashed_passwod"] = user.Hashed_password
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := repo.Collection.UpdateOne(ctx, idfilter(userId), bson.D{{Key: "$set", Value: updateObj}})
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

func (repo *userRepository) SetIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "phone_number", Value: 1}},
	}
	opts := options.CreateIndexes().SetMaxTime(time.Second * 2)
	repo.Collection.Indexes().CreateOne(ctx, indexModel, opts)
}

func idfilter(id string) bson.M {
	var filter = bson.M{}
	_id, _ := primitive.ObjectIDFromHex(id)
	filter = bson.M{"_id": _id}
	return filter
}
