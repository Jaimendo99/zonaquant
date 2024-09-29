package pagehandlers

import (
	"zonaquant/models/uimodels"
	"zonaquant/ui/layout"
	"zonaquant/ui/pages"

	"github.com/labstack/echo/v4"
)

func GetAdminPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Hx-Request") == "" {
			return layout.MainLayout("Admin", pages.AdminPage()).Render(c.Request().Context(), c.Response().Writer)
		}
		return pages.AdminPage().Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUserPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Hx-Request") == "" {
			return layout.MainLayout("User", pages.UserPage()).Render(c.Request().Context(), c.Response().Writer)
		}
		return pages.UserPage().Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetWorkerPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Hx-Request") == "" {
			return layout.MainLayout("Worker", pages.WorkerPage()).Render(c.Request().Context(), c.Response().Writer)
		}
		return pages.WorkerPage().Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetShellPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Hx-Request") == "" {
			return layout.MainLayout("Shell", layout.ApplicationShell(GetAdminSideBarItems())).Render(c.Request().Context(), c.Response().Writer)
		}
		return layout.ApplicationShell(GetAdminSideBarItems()).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetAdminSideBarItems() []uimodels.SideBatItemStruct {
	return []uimodels.SideBatItemStruct{
		{
			ID:    0,
			Title: "Dashboard",
			Icon:  "home",
			Link:  "/admin",
		},
		{
			ID:    1,
			Title: "User",
			Icon:  "person",
			Link:  "/users",
		},
		{
			ID:    2,
			Title: "Worker",
			Icon:  "public",
			Link:  "/worker",
		},
	}
}
