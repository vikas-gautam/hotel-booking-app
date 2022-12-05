package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, fileName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+fileName, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing html templates", err)
		return
	}
}
