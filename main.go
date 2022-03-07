package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("api/legueresponse.json")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(json.Valid(data))

}


