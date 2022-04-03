package repository

import (
	"context"
	"time"

	"github/namuethopro/dobet-user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *userRepository {
	repo := &userRepository{
		Collection: collection,
	}
	repo.SetIndexes()
	return repo
}

func (repo *userRepository) GetUsers(startIndex, perpage int64) ([]domain.User, error) {
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

	err := repo.Collection.FindOne(ctx, bson.D{{Key: "user_id", Value: userId}}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) UpdateUser(userId string, user domain.User) error {
	updateObj := bson.M{}
	switch {
	case len(user.User_id) > 0:
		updateObj["user_id"] = user.User_id
	case len(user.First_name) > 0:
		updateObj["first_name"] = user.First_name
	case len(user.Last_name) > 0:
		updateObj["last_name"] = user.First_name
	case len(user.Phone_number) > 0:
		updateObj["phone_number"] = user.Phone_number
	case len(user.Hashed_password) > 0:
		updateObj["hashed_passwod"] = user.Hashed_password
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := repo.Collection.UpdateOne(ctx, bson.D{{Key: "user_id", Value: userId}}, bson.D{{Key: "$set", Value: updateObj}})
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
	_, err := repo.Collection.DeleteOne(ctx, bson.D{{Key: "user_id", Value: userid}})
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUserBalance(userId string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user := domain.User{}

	err := repo.Collection.FindOne(ctx, bson.D{{Key: "use_id", Value: userId}}).Decode(&user)
	if err != nil {
		return 0.0, err
	}
	return user.Account_balance, nil
}

func (repo *userRepository) AddMoney(userId string, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user := &domain.User{}
	err := repo.Collection.FindOne(ctx, bson.D{{Key: "user_id", Value: userId}}).Decode(&user)
	if err != nil {
		return err
	}
	user.Account_balance += amount
	_, updateObj, err := bson.MarshalValue(user)
	if err != nil {
		return err
	}
	_, err = repo.Collection.UpdateOne(ctx, bson.D{{Key: "user_id", Value: userId}}, bson.D{{Key: "$set", Value: updateObj}})
	return err
}

func (repo *userRepository) SubtractMoney(userId string, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user := &domain.User{}
	err := repo.Collection.FindOne(ctx, bson.D{{Key: "user_id", Value: userId}}).Decode(&user)
	if err != nil {
		return err
	}
	user.Account_balance -= amount
	_, updateObj, err := bson.MarshalValue(user)
	if err != nil {
		return err
	}
	_, err = repo.Collection.UpdateOne(ctx, bson.D{{Key: "user_id", Value: userId}}, bson.D{{Key: "$set", Value: updateObj}})
	return err
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
