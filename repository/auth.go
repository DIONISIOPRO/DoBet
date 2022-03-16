package repository

import (
	"context"
	"errors"
	"time"

	"gitthub.com/dionisiopro/dobet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	Login(phone string) (models.User, error)
	SignUp(user models.User) error
	GetRefreshToken(userId string) (string, error)
	RevokeRefreshToken(userId string) bool
	UpdateRefreshToken(refreshToken, userId string) bool
}

type authRepository struct {
	Collection *mongo.Collection
}

func NewAuthRepository(Collection *mongo.Collection) AuthRepository {
	return &authRepository{
		Collection: Collection,
	}
}

func (repo *authRepository) Login(phone string) (models.User, error) {
	user := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{{"phone_number", phone}}
	err :=repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *authRepository) SignUp(user models.User) error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"phone_number", user.Phone_number}}
	countUser, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if countUser != 0 {
		err = errors.New("this phone number exists, pleasse provide other phone number")
		return err
	}
	_, err = repo.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *authRepository) GetRefreshToken(userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	localuser := models.User{}
	filter := bson.D{{"user_id", userId}}

	err := repo.Collection.FindOne(ctx, filter).Decode(&localuser)
	if err != nil {
		return "", err
	}
	return localuser.RefreshToken, nil

}
func (repo *authRepository) UpdateRefreshToken(refreshToken, userId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"user_id", userId}}
	updateObj := bson.E{"refresh_token", refreshToken}

	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}} )
	if err != nil {
		return false
	}
	return true

}

func (repo *authRepository) RevokeRefreshToken(userId string) bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"user_id", userId}}
	updateObj := bson.E{"refresh_token", ""}

	_, err := repo.Collection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}} )
	if err != nil {
		return false
	}
	return true

}
