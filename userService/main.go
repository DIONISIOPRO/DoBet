package main

import (
	"fmt"
	"github/namuethopro/dobet-user/app"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	Host := os.Getenv("APP_HOST")
	Port := os.Getenv("APP_PORT")
	apphost := fmt.Sprintf("%s:%s", Host, Port)
	engine := gin.Default()
	
	app := app.Application{}
	app.Setup(engine)
	app.Run(apphost)
}
