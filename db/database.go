package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client{
	mongo_url := os.Getenv("MONGO_URL")
	if mongo_url == ""{
		mongo_url = "mongo://localhost:27017"
	}
	fmt.Print(mongo_url)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))
	if err != nil{
		log.Fatal("Error connecting to mongo database")
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil{
		log.Fatal("Error connecting to mongo database")
	}

	return client

}



func OpenCollection(collectionName string) *mongo.Collection{
	database := os.Getenv("DATABASE")
	if database == ""{
		database = "DOBET"
	}
	Client := DbInstance()
	var collection = Client.Database(database).Collection(collectionName)
	return collection
}