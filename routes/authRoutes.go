package routes

import (
	authhandlers "zonaquant/handlers/authHandlers"

	"github.com/labstack/echo/v4"
)

func DefineAuthRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	e.GET("/login", authhandlers.GetPageHandler())
	e.POST("/login", authhandlers.LoginActionHandler())
	e.POST("/signup", authhandlers.CreateClientAccountHandler(), m[0], m[1])
}
