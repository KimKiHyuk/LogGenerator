package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var loggerInstance *log.Logger

func main() {

	loggerInstance = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	loggerInstance.Print("Log service started")

	http.Handle("/healthcheck", new(helloworldHandler))
	http.ListenAndServe(":8080", nil)
	loggerInstance.Print(time.Now()) // logging by 3 sec
}

type helloworldHandler struct {
	http.Handler
}

func (h *helloworldHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	loggerInstance.Print("HTTP Request reached")
	message := "HelloWorld!\n"
	w.Write([]byte(message))
}
