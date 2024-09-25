package authhandlers

import (
	"zonaquant/config"
	dbmodels "zonaquant/models/dbmodels"
	networkmodels "zonaquant/models/networkModels"
	"zonaquant/ui/layout"

	"github.com/labstack/echo/v4"
)

// TODO sign up handler

func CreateClientAccountHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		createUserPayload := networkmodels.CreateUserPayload{}
		err := c.Bind(&createUserPayload)
		if err != nil || verifyCreateUserPayload(createUserPayload) {
			c.Logger().Error(err)
			return layout.FailAlert("Datos no validos").Render(c.Request().Context(), c.Response().Writer)
		}

		db := config.DB

		user := createUserPayload.ToDb()

		user.Accounts = []dbmodels.Account{
			{
				AccountTypeID: 1,
				Balance:       0,
			},
		}

		result := db.Create(&user)
		if result.Error != nil {
			c.Logger().Error(result.Error.Error())
			return layout.FailAlert(result.Error.Error()).Render(c.Request().Context(), c.Response().Writer)
		}

		c.Logger().Info("User created")
		return layout.SuccessAlert("Usuario creado correctamente").Render(c.Request().Context(), c.Response().Writer)
	}
}

func verifyCreateUserPayload(payload networkmodels.CreateUserPayload) bool {
	return payload.Username == "" || payload.Password == "" || payload.RoleID == 0
}
