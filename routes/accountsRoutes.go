package routes

import (
	accounts "zonaquant/handlers/accountsHandlers"

	"github.com/labstack/echo/v4"
)

func DefineAccountsRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	e.GET("/users", accounts.GetUsersHandler())
}
