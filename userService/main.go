package main

import (
	"fmt"
	"github/namuethopro/dobet-user/app"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	Host := os.Getenv("APP_HOST")
	Port := os.Getenv("APP_PORT")
	address := fmt.Sprintf("%s:%s", Host, Port)
	app := app.CreateGinServer()
	app.Run(address)
}
