package main

import (
	"zonaquant/ui"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return ui.Index().Render(c.Request().Context(), c.Response().Writer)
	})
	e.Start(":8080")

}
