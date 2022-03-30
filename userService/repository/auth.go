package repository

import (
	"context"
	"errors"
	"time"

	"github/namuethopro/dobet-user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	Login(phone string) (domain.User, error)
	SignUp(user domain.User) error
	GetRefreshTokens(userId string) ([]string, error)
	AddRefreshToken(refreshToken, userId string) error
	RevokeRefreshToken(userId string) error
	CleanDataBase() error
}

type authRepository struct {
	Collection *mongo.Collection
}

func NewAuthRepository(UserCollection *mongo.Collection) AuthRepository {
	return &authRepository{
		Collection: UserCollection,
	}
}
func (repo *authRepository) CleanDataBase() error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := repo.Collection.DeleteMany(ctx, bson.D{{}})
	return err
}
func (repo *authRepository) Login(phone string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.D{bson.E{Key:"phone_number", Value:phone}}
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *authRepository) SignUp(user domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"phone_number": user.Phone_number}
	countUser, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if countUser != 0 {
		err = errors.New("this phone number exists, pleasse provide other phone number")
		return err
	}
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	_, err = repo.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *authRepository) GetRefreshTokens(userId string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := domain.User{}
	err := repo.Collection.FindOne(ctx, idfilter(userId)).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.RefreshTokens, nil
}
func (repo *authRepository) AddRefreshToken(refreshToken, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := domain.User{}
	err := repo.Collection.FindOne(ctx, idfilter(userId)).Decode(&user)
	if err != nil{
		return err
	}
	user.RefreshTokens = append(user.RefreshTokens, refreshToken)
	_, err = repo.Collection.UpdateOne(ctx, idfilter(userId), bson.D{primitive.E{Key: "$set", Value: user}})
	if err != nil{
		return err
	}
	return nil
}

func (repo *authRepository) RevokeRefreshToken(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	user := domain.User{}
	err := repo.Collection.FindOne(ctx, idfilter(userId)).Decode(&user)
	if err != nil{
		return err
	}
	user.RefreshTokens = []string{}
	_, err = repo.Collection.UpdateOne(ctx, idfilter(userId), bson.D{primitive.E{Key: "$set", Value: user}})
	if err != nil {
		return err
	}
	return nil
}
