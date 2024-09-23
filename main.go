package main

import (
	"fmt"
	"zonaquant/config"
	"zonaquant/routes"
	"zonaquant/ui"

	"github.com/labstack/echo/v4"
)

func init() {
	config.InitDB()
	err := config.Migrate(config.DB)
	if err != nil {
		fmt.Println("error migrating:" + err.Error())
		panic(err)
	}
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return ui.Index().Render(c.Request().Context(), c.Response().Writer)
	})

	routes.DefineAuthRoutes(e)

	e.Static("/static", "static")

	e.Start("0.0.0.0:8080")

}
