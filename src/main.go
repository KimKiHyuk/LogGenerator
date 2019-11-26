package main

import (
	"log"
	"os"
	"time"
)

var loggerInstance *log.Logger

func main() {

	loggerInstance = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	loggerInstance.Print("Log service started")

	for {
		go func(url string) {
			loggerInstance.Print("Fetching from " + url)
		}("https://test.com")

		loggerInstance.Print(time.Now())
		time.Sleep(3 * time.Second) // logging by 3 sec
	}
}
