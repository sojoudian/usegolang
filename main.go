package main

import (
	"fmt"
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

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	switch r.URL.Path {
//	case "/":
//		homehandler(w, r)
//	case "/maz":
//		maz(w, r)
//	default:
//		http.Error(w, "Page note found", http.StatusNotFound)
//		// this is equal to below two lines of code
//		//w.WriteHeader(http.StatusNotFound)
//		//fmt.Fprintf(w, "Page not found")
//	}
//	fmt.Fprintln(w, r.URL.RawPath)
//	fmt.Fprintln(w, r.URL.Path)
//}

type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homehandler(w, r)
	case "/maz":
		maz(w, r)
	case "/faq":
		faq(w, r)
	default:
		http.Error(w, "Page note found", http.StatusNotFound)
	}
}

//func main() {
//	var router http.HandlerFunc
//	router = pathHandler
//	fmt.Println("Starting the server on :3000...")
//	http.ListenAndServe(":3000", router)
//}

//func main() {
//	var router Router
//	//http.HandleFunc("/maz", pathHandler)
//	//http.HandleFunc("/", pathHandler)
//	fmt.Println("starting the server on :3000")
//	http.ListenAndServe(":3000", router)
//
//}

//type Server struct {
//	DB string
//}
//
//func (s *Server) AboutHandler(w http.ResponseWriter, r *http.Request) {
//
//}

func main() {
	//var s Server
	//http.HandleFunc("/about", s.AboutHandler)
	//http.HandleFunc("/", http.HandlerFunc(homehandler).ServeHTTP)
	//http.HandleFunc("/maz", http.HandlerFunc(maz).ServeHTTP)
	//http.Handle("/maz", http.HandlerFunc(maz))
	var router Router
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", router)

}
