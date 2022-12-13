package main

import (
	"log"
	"net/http"

	"github.com/vikas-gautam/hotel-booking-app/pkg/config"
	"github.com/vikas-gautam/hotel-booking-app/pkg/handlers"
	"github.com/vikas-gautam/hotel-booking-app/pkg/render"
)

const portNumber = ":7070"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	app.UseCache = true
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	log.Printf("Starting server on port: %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
