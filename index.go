package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"static/templates/experience.html",
	"static/templates/education.html",
))

func renderTemplate(w http.ResponseWriter, templateName string) {
	err := templates.ExecuteTemplate(w, templateName+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func experienceHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "experience")
}

func educationHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "education")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /experience/", experienceHandler)
	mux.HandleFunc("GET /education/", educationHandler)

	// serve front page by default
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("localhost:9990", mux))
}
