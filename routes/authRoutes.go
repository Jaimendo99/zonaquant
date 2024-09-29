package routes

import (
	authhandlers "zonaquant/handlers/authHandlers"

	"github.com/labstack/echo/v4"
)

// m[0] = valid token, m[1] = admin, m[2] = user, m[3] = worker

func DefineAuthRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	e.GET("/login", authhandlers.GetPageHandler())
	e.POST("/login", authhandlers.LoginActionHandler())
	e.POST("/signup", authhandlers.CreateClientAccountHandler())
}
