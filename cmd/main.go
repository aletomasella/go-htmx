package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(write io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(write, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

type IndexData struct {
	Title string
	Count int
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = newTemplates()

	indexData := IndexData{
		Title: "Hello, World!",
		Count: 0,
	}

	e.GET("/", func(c echo.Context) error {

		indexData.Count++
		return c.Render(200, "index", indexData)

	})

	e.Logger.Fatal(e.Start(":8080"))
}
