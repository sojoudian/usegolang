package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HTMLTpl *template.Template
}

func Parsefile(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template %v", err)
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.HTMLTpl.Execute(w, data) // change nil to user data
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
}
