package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	router := pat.New()

	router.Get("/", dynamicMiddleware.ThenFunc(app.home))
	router.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	router.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.create))
	router.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(router)
}
