package main

import (
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":9090"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting server on port: %s", portNumber)
	http.ListenAndServe(portNumber, nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, fileName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + fileName)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing html templates", err)
	}
}
