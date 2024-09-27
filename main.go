package main

import (
	"fmt"
	"zonaquant/config"
	"zonaquant/routes"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	auths := []echo.MiddlewareFunc{
		echojwt.WithConfig(config.JwtConfig),
		config.AuthorizationMiddleware(1),
		config.AuthorizationMiddleware(2),
		config.AuthorizationMiddleware(3),
	}

	e.Use(middleware.Logger())
	routes.DefineAuthRoutes(e, auths...)
	routes.DefineAccountsRoutes(e, auths...)
	routes.DefinePageRoutes(e, auths...)
	e.Static("/static", "static")

	e.Start("0.0.0.0:8080")

}
