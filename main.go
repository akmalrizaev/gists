package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Gists"))
}

// Add a gistView hadler function
func gistView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific gist..."))
}

// Add a gistCreate handler
func gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new gist"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home) // Restrict this route to exact match / only
	mux.HandleFunc("/gist/view", gistView)
	mux.HandleFunc("/gist/create", gistCreate)

	log.Print(("starting server on :4000"))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
