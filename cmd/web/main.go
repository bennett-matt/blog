package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cache, err := NewTemplateRender()
	if err != nil {
		log.Fatal(err)
	}
	e.Renderer = cache
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, Template{View: "hello.html", Layout: "base"})
	})
	e.Logger.Fatal(e.Start(":1234"))
}
