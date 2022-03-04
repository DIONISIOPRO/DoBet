package main

import (
	"fmt"
	"os"

	//"github.com/gin-gonic/gin"

	// "log"
	// "math"
	// "net/http"
	// "strconv"
	// "time"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/database"
	// 	"github.com/go-playground/validator/v10"
	// 	"go.mongodb.org/mongo-driver/bson"
	// 	"go.mongodb.org/mongo-driver/bson/primitive"
	// 	"go.mongodb.org/mongo-driver/mongo"
	// 	"go.mongodb.org/mongo-driver/mongo/options"
	//
)

func main() {
	var port = os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	route := gin.New()

	route.Use(gin.Logger())

	var collection = database.OpenCollection("users")
	fmt.Print(collection)

	route.Run(port)

	//fmt.Print("Running")

}
