package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	tomorow := now + int64(time.Hour * 24)
	fmt.Print(tomorow - now)
}