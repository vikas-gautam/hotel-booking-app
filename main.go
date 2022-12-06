package main

import (
	"log"
	"net/http"

	"github.com/vikas-gautam/hotel-booking-app/pkg/handlers"
)

const portNumber = ":7070"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting server on port: %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}
