package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", mainHandler)
	// serve static files (look inside 'static' directory)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Serving on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
