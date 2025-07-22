package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", app.home)                // Restrict this route to exact match / only
	mux.HandleFunc("GET /gist/view/{id}", app.gistView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /gist/create", app.gistCreate)
	mux.HandleFunc("POST /gist/create", app.gistCreatePost)

	// log.Printf("starting server on %s", *addr)
	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
