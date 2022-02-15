package main

import (
	"os"

	//"github.com/gin-gonic/gin"

	// "log"
	// "math"
	// "net/http"
	// "strconv"
	// "time"

	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 
)

func main() {
	var port = os.Getenv("PORT")

	if port == ""{
		port = ":8080"
	}

	route := gin.New()

	route.Use(gin.Logger())

	route.Run(port)
	


}