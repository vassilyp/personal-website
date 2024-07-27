package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am here")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<p>world</p>")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/", helloHandler)

	// serve front page by default
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
	// http.ListenAndServe("localhost:8080", nil)
}
