package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// w.Write([]byte("Hello from Gists"))
}

// Add a gistView hadler function
func gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// msg := fmt.Sprintf("Display a specific gist with ID %d...", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific gist with ID %d...", id)
}

// Add a gistCreate handler
func gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new gist"))
}

// Add a gistCreatePost handler function
func gistCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new gist..."))
}
