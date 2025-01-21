package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/about.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/contact.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/snippetView.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", id)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not Allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
