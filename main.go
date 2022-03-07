package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitthub.com/dionisiopro/dobet/models"
)

func main() {

	data, err := ioutil.ReadFile("api/legueresponse.json")
	if err != nil {
		fmt.Print(err)
	}

	var league models.League

	err = json.Unmarshal(data, &league)
	if err != nil {
		fmt.Errorf("error %v", err)
	}

	fmt.Println(league)

}
