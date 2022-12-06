package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func RenderTemplate(w http.ResponseWriter, fileName string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+fileName, "./templates/base.layout.html")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		log.Println("error parsing html templates", err)
// 		return
// 	}
// }

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {

	var template *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached templates")
	}

	template = tc[t]
	err = template.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	temp, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = temp

	return nil
}
