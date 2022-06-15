package controllers

import (
	"github.com/sojoudian/usegolang/views"
	"net/http"
)

func StatickHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
