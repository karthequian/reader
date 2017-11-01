package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	runs := 1
	for runs > 0 {
		fmt.Printf("Log level passed via env variables was: '%v'\n", os.Getenv("log_level"))
		time.Sleep(time.Second * 5)
	}
}
