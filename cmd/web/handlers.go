package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Gists"))
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
