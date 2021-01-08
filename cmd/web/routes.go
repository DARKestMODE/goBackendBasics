package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := pat.New()

	router.Get("/", http.HandlerFunc(app.home))
	router.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	router.Post("/snippet/create", http.HandlerFunc(app.create))
	router.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
