package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		jsonRepose := make(map[string]interface{})

		jsonRepose["message"] = "Hello, World!"

		return c.JSON(200, jsonRepose)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
