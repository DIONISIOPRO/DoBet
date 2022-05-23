package repository

import (
	"context"
	"errors"
	"github.com/dionisiopro/dobet-auth/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	Collection *mongo.Collection
}

func NewAuthRepository(UserCollection *mongo.Collection) *authRepository {
	return &authRepository{
		Collection: UserCollection,
	}
}

func (repo *authRepository) AddUser(user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	doc := prepareUserToSave(user)
	_, err := repo.Collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}

func (repo *authRepository) UpdateUser(userid string, user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	doc := prepareUserToSave(user)
	filter := bson.D{{Key: "user_id", Value: userid}}
	result := repo.Collection.FindOneAndUpdate(ctx, filter, doc)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (repo *authRepository) RemoveUser(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{Key: "user_id", Value: userId}}
	result := repo.Collection.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (repo *authRepository) Login(phone string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{bson.E{Key: "phone_number", Value: phone}}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (repo *authRepository) AddRefreshToken(phone_number, refreshtoken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := domain.User{}
	err := repo.Collection.FindOne(ctx, bson.D{{Key: "phone_number", Value: phone_number}}).Decode(&user)
	if err != nil {
		return err
	}
	user.RefreshTokens = append(user.RefreshTokens, refreshtoken)
	doc := prepareUserToSave(user)
	_, err = repo.Collection.UpdateOne(ctx, bson.D{{Key: "phone_number", Value: phone_number}}, doc)
	if err != nil {
		return err
	}
	return nil
}

func (repo *authRepository) GetRefreshTokens(userid string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := domain.User{}
	err := repo.Collection.FindOne(ctx, bson.D{{Key: "user_id", Value: userid}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.RefreshTokens, nil
}

func (repo *authRepository) SignUp(user domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"phone_number": user.Phone_number}
	countUser, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return "", err
	}
	if countUser != 0 {
		err = errors.New("this phone number exists, pleasse provide other phone number")
		return "", err
	}
	userDoc := prepareUserToSave(user)
	_, err = repo.Collection.InsertOne(ctx, userDoc)
	if err != nil {
		return "", err
	}
	return userDoc.Map()["user_id"].(string), nil
}

func prepareUserToSave(user domain.User) bson.D {
	id := primitive.NewObjectID()
	_id := bson.E{Key: "_d", Value: id}
	userId := bson.E{Key: "user_id", Value: id.Hex()}
	firstName := bson.E{Key: "first_name", Value: user.First_name}
	lastName := bson.E{Key: "last_name", Value: user.Last_name}
	phone := bson.E{Key: "phone_number", Value: user.Phone_number}
	refreshToken := bson.E{Key: "refresh_tokens", Value: []string{}}
	hash := bson.E{Key: "hashed_password", Value: user.Hashed_password}
	doc := bson.D{_id, userId, firstName,
		lastName, phone, hash, refreshToken,
	}
	return doc
}
