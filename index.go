package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

var templates = template.Must(template.ParseFiles(
	"static/templates/experience.html",
	"static/templates/education.html",
	"static/templates/about.html",
	"static/templates/projects.html",
	"static/glsl/vert.glsl",
	"static/glsl/frag.glsl",
))

func renderTemplate(w http.ResponseWriter, templateName string) {
	err := templates.ExecuteTemplate(w, templateName, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /experience/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "experience.html")
	})
	mux.HandleFunc("GET /education/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "education.html")
	})
	mux.HandleFunc("GET /about/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about.html")
	})
	mux.HandleFunc("GET /projects/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "projects.html")
	})
	mux.HandleFunc("GET /static/glsl/{filename}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello", r.PathValue("filename"))
		renderTemplate(w, r.PathValue("filename"))
	})
	
	// serve front page by default
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("localhost:9990", mux))
}
