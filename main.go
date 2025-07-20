package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Gists"))
}

// Add a gistView hadler function
func gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific gist with ID %d...", id)
	w.Write([]byte(msg))
}

// Add a gistCreate handler
func gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new gist"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)                // Restrict this route to exact match / only
	mux.HandleFunc("/gist/view/{id}", gistView) // Add the {id} wildcard segment
	mux.HandleFunc("/gist/create", gistCreate)

	log.Print(("starting server on :4000"))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
