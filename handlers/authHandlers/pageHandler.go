package authhandlers

import (
	authui "zonaquant/ui/auth"
	layoutui "zonaquant/ui/layout"

	"github.com/labstack/echo/v4"
)

func GetPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return layoutui.MainLayout("Login", authui.LoginComp()).Render(c.Request().Context(), c.Response().Writer)
	}
}
