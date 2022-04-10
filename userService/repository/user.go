package repository

import (
	"context"
	"errors"
	"github/namuethopro/dobet-user/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type MongoDriverUserCollection interface {
	CountDocuments(ctx context.Context, filter interface{},
		opts ...*options.CountOptions) (int64, error)
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type userRepository struct {
	Collection MongoDriverUserCollection
}

func NewUserRepository(collection MongoDriverUserCollection) *userRepository {
	repo := &userRepository{
		Collection: collection,
	}
	return repo
}

func (repo *userRepository) CreateUser(user domain.User) (string, error) {
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

func (repo *userRepository) GetUsers(startIndex, perpage int64) ([]domain.User, error) {
	users := []domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	opts := options.Find()
	opts.Limit = &perpage
	opts.Skip = &startIndex
	cursor, err := repo.Collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
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
	if userid == ""{
		return errors.New("invalid params")
	}
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
		return float64(-1), err
	}
	return user.Account_balance, nil
}

func (repo *userRepository) AddMoney(userId string, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user := &domain.User{}
	result := repo.Collection.FindOne(ctx, bson.D{{Key: "user_id", Value: userId}})
	if result == nil{
		return errors.New("user id does not exist")
	}
	err := result.Decode(&user)
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

func prepareUserToSave(user domain.User) bson.D {
	user.Hashed_password = hasFromPassword(user.Password)
	id := primitive.NewObjectID()
	_id := bson.E{Key: "_d", Value: id}
	userId := bson.E{Key: "user_id", Value: id.Hex()}
	firstName := bson.E{Key: "first_name", Value: user.First_name}
	lastName := bson.E{Key: "last_name", Value: user.Last_name}
	phone := bson.E{Key: "phone_number", Value: user.Phone_number}
	balance := bson.E{Key: "account_balance", Value: user.Account_balance}
	created := bson.E{Key: "created_at", Value: time.Now().Local()}
	update := bson.E{Key: "updated_at", Value: time.Now().Local()}
	isAdmin := bson.E{Key: "is_admin", Value: user.IsAdmin}
	refreshToken := bson.E{Key: "refresh_tokens", Value: []string{}}
	hash := bson.E{Key: "hashed_password", Value: user.Hashed_password}
	doc := bson.D{_id, userId, firstName,
		lastName, phone, hash,
		balance, isAdmin, refreshToken,
		created, update,
	}
	return doc
}
func hasFromPassword(password string) string {
	data, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(data)
}
