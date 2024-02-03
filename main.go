package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var fileServer http.Handler = http.FileServer(http.Dir("./static"))
	var logger *log.Logger = log.New(os.Stdout, time.Now().String()+" :: INFO :: ", 1)

	http.Handle("/", fileServer)
	logger.Printf("Starting Http Server")
	var err error = http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("Server startup failed !!")
	}
	logger.Printf("Stopping Http Server")

}
