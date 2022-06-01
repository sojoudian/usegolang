package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func maz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Maz</h1>")
}

func main() {
	http.HandleFunc("/maz", maz)
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting the server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
