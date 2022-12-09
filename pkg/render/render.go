package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vikas-gautam/hotel-booking-app/pkg/config"
)

//render template advance

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	fmt.Println("executing newTemplate function")
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	fmt.Println("executing renderTemplate func")

	// No need to create template cache, already created during application run
	tc := app.TemplateCache

	// get requested template from cache
	fmt.Println("requesting template from cache")

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("error while requesting tmpl from cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	fmt.Println("executing CreateTemplateCache")

	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	fmt.Println(pages)
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
