package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	tc *TemplateCache
}

func main() {
	templateCache, err := NewTemplateRender()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	app := &application{tc: templateCache}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		app.render(w, r, Template{View: "hello.html", Layout: "base"})
	})

	r.Handle("/dist/*", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	http.ListenAndServe(":1234", r)
}

func (a *application) render(w http.ResponseWriter, r *http.Request, t Template) {
	template, ok := a.tc.templates[t.View]
	if !ok {
		// TODO: handle this
	}

	template.ExecuteTemplate(w, t.Layout, t.Data)
}
