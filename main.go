package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sojoudian/usegolang/controllers"
	"github.com/sojoudian/usegolang/views"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type User struct {
	HostName string
	Option_a string
	Option_b string
	Vote     string
}

var (
	HomeTemplate views.Template
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	t, err := views.Parsefile(filePath)
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
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

func help(w http.ResponseWriter, r *http.Request) {
	tpl := filepath.Join("templates", "help.gohtml")
	executeTemplate(w, tpl)
}

func s4e(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	type Fav struct {
		Movie string
		Music string
		Pi    float64
	}
	type S4E struct {
		Name   string
		Bio    string
		Age    int
		Allfav Fav
	}
	userInfo := S4E{
		Name: "Maz",
		Bio:  "this is the bio",
		Age:  32,
		Allfav: Fav{
			Movie: "Batman",
			//Music: "Metallica",
			Pi: 3.14,
		},
	}
	tpl, err := template.ParseFiles("templates/s4e.gohtml")
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

func faq(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parsefile(filepath.Join("templates", "index.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/home", controllers.StatickHandler(tpl))

	r.Get("/", homehandler) // duplicate but I need it for the future use
	r.Get("/faq", faq)

	tpl, err = views.Parsefile(filepath.Join("templates", "help.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/help", controllers.StatickHandler(tpl))
	//r.Get("/help", help)

	tpl, err = views.Parsefile(filepath.Join("templates", "s4e.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StatickHandler(tpl))
	//r.Get("/s4e", s4e)

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found!!!!", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", r)

}
