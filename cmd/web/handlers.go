package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		// log.Print(err.Error())
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// log.Print(err.Error())
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
	}

	// w.Write([]byte("Hello from Gists"))
}

// Add a gistView hadler function
func (app *application) gistView(w http.ResponseWriter, r *http.Request) {
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
func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new gist"))
}

// Add a gistCreatePost handler function
func (app *application) gistCreatePost(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.gists.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/gist/view/%d", id), http.StatusSeeOther)

}
