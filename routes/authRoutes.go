package routes

import (
	authhandlers "zonaquant/handlers/authHandlers"

	"github.com/labstack/echo/v4"
)

func DefineAuthRoutes(e *echo.Echo) {
	e.GET("/login", authhandlers.GetPageHandler())
	e.POST("/signup", authhandlers.CreateClientAccountHandler())
}
