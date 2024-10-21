package routes

import (
	pagehandlers "zonaquant/handlers/pageHandlers"

	"github.com/labstack/echo/v4"
)

func DefinePageRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	// Define routes here
	e.GET("/admin", pagehandlers.GetAdminPage(), m[0], m[1])
	e.GET("/user", pagehandlers.GetUserPage(), m[0], m[2])
	e.GET("/worker", pagehandlers.GetWorkerPage(), m[0], m[3])
	e.GET("/", pagehandlers.GetShellPage())
}
