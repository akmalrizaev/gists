package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", home)                // Restrict this route to exact match / only
	mux.HandleFunc("GET /gist/view/{id}", gistView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /gist/create", gistCreate)
	mux.HandleFunc("POST /gist/create", gistCreatePost)

	// log.Printf("starting server on %s", *addr)
	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
