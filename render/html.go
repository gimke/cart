package render

import (
	"html/template"
	"net/http"
)

type (
	Delims struct {
		Left  string
		Right string
	}

	HTML struct {
		Template *template.Template
		Name     string
		Data     interface{}
	}
)

var htmlContentType = []string{"text/html; charset=utf-8"}

func (r HTML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	if len(r.Name) == 0 {
		return r.Template.Execute(w, r.Data)
	}
	return r.Template.ExecuteTemplate(w, r.Name, r.Data)
}

func (r HTML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, htmlContentType)
}
