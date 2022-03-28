package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gitthub.com/dionisiopro/dobet/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadConfig() config.BaseConfig {
	Config := config.BaseConfig{}

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileDecoder := json.NewDecoder(file)
	err = fileDecoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
	return Config
}
func DbInstance() *mongo.Client {
	config := LoadConfig()
	url := fmt.Sprintf("%v:%v",config.DB.Host, config.DB.Port)

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
