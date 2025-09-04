package main

import (
	htmltemplate "html/template"
	"log"
	"net/http"
	texttemplate "text/template"
)

var htmltemplates = htmltemplate.Must(htmltemplate.ParseFiles(
	"static/templates/experience.html",
	"static/templates/education.html",
	"static/templates/about.html",
	"static/templates/projects.html",
))

var texttemplates = texttemplate.Must(texttemplate.ParseFiles(
	"static/glsl/vert.glsl",
	"static/glsl/frag.glsl",
))

func renderHTMLTemplate(w http.ResponseWriter, templateName string) {
	err := htmltemplates.ExecuteTemplate(w, templateName, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderTextTemplate(w http.ResponseWriter, templateName string) {
	err := texttemplates.ExecuteTemplate(w, templateName, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /experience/", func(w http.ResponseWriter, r *http.Request) {
		renderHTMLTemplate(w, "experience.html")
	})
	mux.HandleFunc("GET /education/", func(w http.ResponseWriter, r *http.Request) {
		renderHTMLTemplate(w, "education.html")
	})
	mux.HandleFunc("GET /about/", func(w http.ResponseWriter, r *http.Request) {
		renderHTMLTemplate(w, "about.html")
	})
	mux.HandleFunc("GET /projects/", func(w http.ResponseWriter, r *http.Request) {
		renderHTMLTemplate(w, "projects.html")
	})
	mux.HandleFunc("GET /static/glsl/{filename}", func(w http.ResponseWriter, r *http.Request) {
		renderTextTemplate(w, r.PathValue("filename"))
	})

	// serve front page by default
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("localhost:9990", mux))
}
