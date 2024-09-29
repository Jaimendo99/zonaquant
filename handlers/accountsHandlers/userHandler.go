package accounts

import (
	"zonaquant/config"
	models "zonaquant/models/dbmodels"
	"zonaquant/models/uimodels"
	"zonaquant/ui/pages"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := config.DB
		users := []models.User{}
		db.Preload("Role").Preload("Accounts.AccountType").Find(&users)
		uiusers := []uimodels.UIuser{}
		for _, user := range users {
			uiusers = append(uiusers, uimodels.UserToUI(&user))
		}
		return pages.UserListGrid(uiusers).Render(c.Request().Context(), c.Response().Writer)
	}
}
