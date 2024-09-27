package accounts

import (
	"zonaquant/config"
	models "zonaquant/models/dbmodels"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO verify if user is authenticated

		db := config.DB
		users := []models.User{}
		db.Find(&users)
		return c.JSON(200, users)
	}
}
