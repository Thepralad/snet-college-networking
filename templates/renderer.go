package render

import (
	"html/template"
	"net/http"
)

func RenderTemplate(res http.ResponseWriter, tmpl string, data any) error {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		return err
	}
	t.Execute(res, data)
	return nil
}
