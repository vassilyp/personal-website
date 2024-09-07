package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"static/templates/experience.html",
	"static/templates/education.html",
	"static/templates/about.html",
	"static/templates/projects.html",
))

func renderTemplate(w http.ResponseWriter, templateName string) {
	err := templates.ExecuteTemplate(w, templateName+".html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /experience/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "experience")
	})
	mux.HandleFunc("GET /education/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "education")
	})
	mux.HandleFunc("GET /about/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about")
	})
	mux.HandleFunc("GET /projects/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "projects")
	})

	// serve front page by default
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("localhost:9990", mux))
}
