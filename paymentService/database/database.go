package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"

)

func DbInstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil{
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
	config := LoadConfig()
	Client := DbInstance()
	db := config.DB.DB

	var collection = Client.Database(db).Collection(collectionName)
	return collection
}
