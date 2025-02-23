package render

import (
	"html/template"
	"net/http"
)

func RenderTemplate(res http.ResponseWriter, tmpl string, data any) {
	t, _ := template.ParseFiles("./templates/" + tmpl + ".html")
	t.Execute(res, data)
}
