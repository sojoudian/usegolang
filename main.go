package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	HostName string
	Option_a string
	Option_b string
	Vote     string
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)

	userInfo := User{
		HostName: hostname,
		Option_a: "Cats",
		Option_b: "Dogs",
		Vote:     "",
	}
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, userInfo)
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
}

func maz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello from Maz</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Faq page</h1>
<ul>
	<li>first</li>
</ul>
`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homehandler)
	r.Get("/faq", faq)
	r.Get("/maz", maz)
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found!!!!", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", r)

}
