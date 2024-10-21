package authhandlers

import (
	"zonaquant/ui/authui"
	"zonaquant/ui/layout"
	"zonaquant/ui/pages"

	"github.com/labstack/echo/v4"
)

func GetPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return layout.MainLayout("Login", authui.LoginComp()).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetRegisterPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return pages.AddUserForm().Render(c.Request().Context(), c.Response().Writer)
	}
}
