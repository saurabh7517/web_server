package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger *log.Logger = log.New(os.Stdout, time.Now().String()+" :: INFO :: ", 1)

func handleForm(w http.ResponseWriter, r *http.Request) {
	var err error = r.ParseForm()
	if err != nil {
		logger.Println(err, "cannot parse form")
		return
	}
	logger.Println("Form parsed successfully")
	var firstname string = r.FormValue("fname")
	var lastname string = r.FormValue("lname")
	logger.Printf("First Name :: %s", firstname)
	logger.Printf("Last Name :: %s", lastname)
}

func main() {
	var fileServer http.Handler = http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)

	logger.Printf("Starting Http Server")
	var err error = http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("Server startup failed !!")
	}
	logger.Printf("Stopping Http Server") // this needs to be implemented to tackle gracefully closing of servers

}
