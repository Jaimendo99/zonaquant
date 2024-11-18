package accounts

import (
	"net/http"
	"zonaquant/config"
	models "zonaquant/models/dbmodels"
	"zonaquant/models/uimodels"
	"zonaquant/ui/layout"
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

func DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := config.DB
		id := c.Param("id")
		user := models.User{}
		db.Preload("Accounts").First(&user, id)
		for _, account := range user.Accounts {
			if account.Balance != 0 {
				c.Response().WriteHeader(http.StatusBadRequest)
				return layout.SimpleNotification("El usuario tiene cuentas con saldo", true, nil).Render(c.Request().Context(), c.Response().Writer)
			}
			result := db.Unscoped().Delete(&account)
			if result.Error != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				return layout.SimpleNotification(result.Error.Error(), true, nil).Render(c.Request().Context(), c.Response().Writer)
			}
		}
		db.Unscoped().Delete(&user)
		return GetUsersHandler()(c)
	}
}

func ResetPasswordHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := config.DB
		id := c.Param("id")
		user := models.User{}
		db.First(&user, id)
		newPass := genNewPassword()
		user.Password = newPass
		result := db.Save(&user)
		if result.Error != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return layout.SimpleNotification(result.Error.Error(), true, nil).Render(c.Request().Context(), c.Response().Writer)
		}
		component := pages.PasswordReseted(newPass)
		return layout.SimpleNotification("", false, component).Render(c.Request().Context(), c.Response().Writer)
	}
}

func genNewPassword() string {
	return "123456"
}
