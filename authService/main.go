package main

import (
	"fmt"
	"os"
	"github/namuethopro/dobet-user/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	Host := os.Getenv("APP_HOST")
	Port := os.Getenv("APP_PORT")
	apphost := fmt.Sprintf("%s:%s",Host,Port)
	app := app.NewApplication(apphost)
	app.Run()
}