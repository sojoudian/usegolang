package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	bio := `<script>alert("You were hacked")</script>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Index</h1><p>Bio:"+bio+"</p>")
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
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found!!!!", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", r)

}
