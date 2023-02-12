package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vikas-gautam/hotel-booking-app/internal/config"
	"github.com/vikas-gautam/hotel-booking-app/internal/driver"
	"github.com/vikas-gautam/hotel-booking-app/internal/handlers"
	"github.com/vikas-gautam/hotel-booking-app/internal/helpers"
	"github.com/vikas-gautam/hotel-booking-app/internal/models"
	"github.com/vikas-gautam/hotel-booking-app/internal/render"
)

const portNumber = ":7070"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("Starting mail listner")
	ListenForMail()

	log.Printf("Starting server on port: %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

// to make main testable
func run() (*driver.DB, error) {

	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=127.0.0.1 port=5432 dbname=bookings user=postgres password=password")
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	log.Println("Connected to database")

	tc, err := render.CreateTemplateCache()
	// fmt.Printf("checking template cache %v\n", tc)
	if err != nil {
		log.Fatal("cannot create template cache from main")
		return nil, err
	}
	app.TemplateCache = tc

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	app.UseCache = false
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)

	return db, nil
}
