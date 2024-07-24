package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("page/index.html"))
var validPath = regexp.MustCompile("^/$")

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// redirect to main page if given weird URL path
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// TODO: keep in mind we can pass along template variables into this!!!
	templates.Execute(w, "page/index.html")
	return
}

func main() {
	http.HandleFunc("/", mainHandler)

	fmt.Println("Serving on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
