package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

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
	fmt.Print("Connected to the client" + url)

	return client

}

func OpenCollection(collectionName string) *mongo.Collection {
	Client := DbInstance()
	var collection = Client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}
