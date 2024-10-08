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
	// Hice un cambio
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = newTemplates()

	indexData := IndexData{
		Title: "Hello!",
		Count: 0,
	}

	availableAction := make(map[string]bool)

	availableAction["greet"] = true
	availableAction["curse"] = true
	availableAction["compliment"] = true

	e.GET("/", func(c echo.Context) error {

		user := "Random User"

		indexData.Title = "Hello, " + user + "!"
		indexData.Count++
		return c.Render(200, "index", indexData)

	})

	e.GET("/:user", func(c echo.Context) error {
		user := c.Param("user")

		indexData.Title = "Hello, " + user + "!"
		indexData.Count++
		return c.Render(200, "index", indexData)

	})

	e.GET("/:user/:action" func(c echo.Context) error {
		user := c.Param("user")
		action := c.Param("action")

		if _, ok := availableAction[action]; !ok {
			return c.String(400, "Action not available")
		}

		indexData.Title = "Hello, " + user + "!"
		indexData.Count++
		return c.Render(200, "index", indexData)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
