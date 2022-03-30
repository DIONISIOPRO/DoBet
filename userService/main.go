package main

import "github/namuethopro/dobet-user/app"

func main() {
	app := app.NewApplication(":8080")
	app.Run()
}