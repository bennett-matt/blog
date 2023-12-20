package main

import (
	"errors"
	"fmt"
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

type Template struct {
	View   string
	Layout string
	Data   interface{}
}

type TemplateCache struct {
	templates map[string]*template.Template
}

func Render(c echo.Context, status int, t Template) error {
	return c.Render(status, t.View, t)
}

func (tc *TemplateCache) Render(w io.Writer, view string, data interface{}, c echo.Context) error {
	fmt.Println("data: ", data)
	t, ok := data.(Template)
	if !ok {
		return ErrNoTemplate
	}

	fmt.Println("templates: ", tc.templates)
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

	fmt.Println("pages: ", pages)

	return &TemplateCache{templates: cache}, nil
}
