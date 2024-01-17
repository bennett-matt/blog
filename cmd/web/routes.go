package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type envelope map[string]any

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.Handler(http.MethodGet, "/dist/*filepath", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	// TODO: separate not found logic to handlers file
	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		app.render(w, r, Template{View: "hello.html", Layout: "base"})
	})
	router.HandlerFunc(http.MethodGet, "/sessions/new", app.signInHandler)
	router.HandlerFunc(http.MethodGet, "/sessions/reset", app.resetPasswordHandler)
	router.HandlerFunc(http.MethodGet, "/users/new", app.userSignupHandler)

	standard := alice.New()
	return standard.Then(router)
}
