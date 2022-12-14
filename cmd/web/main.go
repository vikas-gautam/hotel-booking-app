package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vikas-gautam/hotel-booking-app/pkg/config"
	"github.com/vikas-gautam/hotel-booking-app/pkg/handlers"
	"github.com/vikas-gautam/hotel-booking-app/pkg/render"
)

const portNumber = ":7070"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
