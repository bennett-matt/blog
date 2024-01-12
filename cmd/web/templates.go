package main

import (
	"errors"
	"html/template"
	"io/fs"
	"path/filepath"

	"io"

	"github.com/bennett-matt/blog/public"
	"github.com/labstack/echo/v4"
)

var (
	functions               = template.FuncMap{}
	ErrTemplateDoesNotExist = errors.New("template doesn't exist")
	ErrNoTemplate           = errors.New("not a template")
)

type TemplateData struct {
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	Data            any
}

type Template struct {
	View   string
	Layout string
	Data   TemplateData
}

type TemplateCache struct {
	templates map[string]*template.Template
}

func (tc *TemplateCache) Render(w io.Writer, view string, data interface{}, c echo.Context) error {
	t, ok := data.(Template)
	if !ok {
		return ErrNoTemplate
	}

	if _, ok := tc.templates[view]; !ok {
		return ErrTemplateDoesNotExist
	}

	return tc.templates[view].ExecuteTemplate(w, t.Layout, t.Data)
}

func NewTemplateRender() (*TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(public.Files, "views/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"layouts/*.html",
			"partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(public.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return &TemplateCache{templates: cache}, nil
}
