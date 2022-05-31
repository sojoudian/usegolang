package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting the server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
