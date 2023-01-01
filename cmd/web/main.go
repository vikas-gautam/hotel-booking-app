package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vikas-gautam/hotel-booking-app/internal/config"
	"github.com/vikas-gautam/hotel-booking-app/internal/handlers"
	"github.com/vikas-gautam/hotel-booking-app/internal/models"
	"github.com/vikas-gautam/hotel-booking-app/internal/render"
)

const portNumber = ":7070"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on port: %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

// to make main testable
func run() error {

	gob.Register(models.Reservation{})

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	fmt.Printf("checking template cache %v\n", tc)
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}
	app.TemplateCache = tc

	app.UseCache = false
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	return nil
}
