package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/akmalrizaev/gists/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Gists = gists

	app.render(w, r, http.StatusOK, "home.tmpl", data)

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// data := templateData{
	// 	Gists: gists,
	// }

	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {

	// 	app.serverError(w, r, err)
	// }

}

// Add a gistView hadler function
func (app *application) gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	gist, err := app.gists.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Gist = gist

	app.render(w, r, http.StatusOK, "view.tmpl", data)

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/view.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// data := templateData{
	// 	Gist: gist,
	// }

	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }

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
