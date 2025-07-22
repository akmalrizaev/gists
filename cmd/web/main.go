package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", home)                // Restrict this route to exact match / only
	mux.HandleFunc("GET /gist/view/{id}", gistView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /gist/create", gistCreate)
	mux.HandleFunc("POST /gist/create", gistCreatePost)

	log.Print(("starting server on :4000"))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
