package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Index</h1>")
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

// HTTP handler accessing the url routing parameters.
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"userID"` from the request of a matching
	// routing pattern. An example routing pattern could be: /users/{userID}
	userID := chi.URLParam(r, "userID")

	// fetch `"key"` from the request context
	ctx := r.Context()
	key := ctx.Value("key").(string)

	// respond to the client
	w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homehandler)
	r.Get("/faq", faq)
	r.Get("/maz", maz)
	r.Get("/users/{userID}", MyRequestHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found!!!!", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", r)

}

//
//type Router struct {
//}
//
//func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homehandler(w, r)
//	case "/maz":
//		maz(w, r)
//	case "/faq":
//		faq(w, r)
//	default:
//		http.Error(w, "Page note found", http.StatusNotFound)
//	}
//}
