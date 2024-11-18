package routes

import (
	accounts "zonaquant/handlers/accountsHandlers"

	"github.com/labstack/echo/v4"
)

func DefineAccountsRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	e.GET("/users", accounts.GetUsersHandler(), m[0], m[1])
	e.DELETE("/users/:id", accounts.DeleteUserHandler(), m[0], m[1])
	e.GET("/resetpass/:id", accounts.ResetPasswordHandler(), m[0], m[1])
}
