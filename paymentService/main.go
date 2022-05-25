package main

import (
	"fmt"
	"os"
	"github.com/dionisiopro/dobet_payment/app"
	"github.com/joho/godotenv"
)

// @title Payment API
// @version 1.0
// @description This is a payment service for DoBet Application
// @termsOfService http://swagger.io/terms/

// @contact.name Dionisio Paulo
// @contact.url meusite.com
// @contact.email dionisiopaulonamuetho@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9003``
// @BasePath api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	done := make(chan bool)
	defer func() {
		done <- true
	}()
	Host := os.Getenv("APP_HOST")
	Port := os.Getenv("APP_PORT")
	address := fmt.Sprintf("%s:%s", Host, Port)
	app := app.CreateGinServer(done)
	app.Run(address)
}
