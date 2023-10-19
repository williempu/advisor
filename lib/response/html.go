package response

import (
	"html/template"
	"log"
	"net/http"
)

// RenderHTML gets the filename and data then parse them into template
func RenderHTML(w http.ResponseWriter, r *http.Request, fileName string, data map[string]interface{}) {

	var files = []string{
		"templates/index.html",
	}

	var templates = template.Must(template.ParseFiles(files...)) // alternative way, use ParseFiles

	// try parsing all files and pass data if any, the root is index.html always (views/index.html)
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
