package main

import (
	"errors"
	"fmt"
	"module1/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.gohtml", &templateData{Snippets: s,})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.gohtml", &templateData{Snippet: s,})

}

func (app *application) create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed methods", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Test"
	content := "test test test"
	expires := 21

	id ,err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
