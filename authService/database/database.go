package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can`t load env file")
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	url := fmt.Sprintf("%v:%v", DB_HOST, DB_PORT)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Error connecting to mongo database")
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to mongo database")
	}
	log.Print("Connected to the client" + url)

	return client

}

func OpenCollection(collectionName string) *mongo.Collection {
	DB_NAME := os.Getenv("DB_NAME")
	Client := DbInstance()
	var collection = Client.Database(DB_NAME).Collection(collectionName)
	return collection
}
