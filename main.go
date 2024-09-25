package main

import (
	"fmt"
	"zonaquant/config"
	"zonaquant/routes"

	"github.com/labstack/echo/v4"
)

func init() {
	config.InitDB()
	err := config.Migrate(config.DB)
	if err != nil {
		fmt.Println("error migrating:" + err.Error())
		panic(err)
	}

	config.LoadEnvVariables()
}

func main() {
	e := echo.New()

	routes.DefineAuthRoutes(e)

	e.Static("/static", "static")

	e.Start("0.0.0.0:8080")

}
