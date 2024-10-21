package authhandlers

import (
	"net/http"
	"strconv"
	"time"
	"zonaquant/config"
	dbmodels "zonaquant/models/dbmodels"
	networkmodels "zonaquant/models/networkModels"
	"zonaquant/ui/layout"
	"zonaquant/ui/pages"

	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// TODO sign up handler

func CreateClientAccountHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		createUserPayload := networkmodels.CreateUserPayload{}
		err := c.Bind(&createUserPayload)
		if verifyCreateUserPayload(createUserPayload) || err != nil {
			c.Logger().Error(err)
			return layout.SimpleNotification("Datos no validos", true).Render(c.Request().Context(), c.Response().Writer)
		}

		db := config.DB
		user := createUserPayload.ToDb()

		user.Accounts = []dbmodels.Account{
			{
				AccountNumber: generateAccountNumber(user, 1),
				AccountTypeID: 1,
				Balance:       0,
			},
		}
		if createUserPayload.Box {
			user.Accounts = append(user.Accounts, dbmodels.Account{
				AccountNumber: generateAccountNumber(user, 4),
				AccountTypeID: 4,
				Balance:       0,
			})
		}

		result := db.Create(&user)
		if result.Error != nil {
			c.Logger().Error(result.Error.Error())
			return layout.SimpleNotification(result.Error.Error(), true).Render(c.Request().Context(), c.Response().Writer)
		}

		c.Logger().Info("User created")
		return layout.SimpleNotification("Usuario creado", false).Render(c.Request().Context(), c.Response().Writer)
	}
}

func LoginActionHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginPayload := networkmodels.LoginPayload{}

		err := c.Bind(&loginPayload)
		if err != nil {
			c.Logger().Error(err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return layout.SimpleNotification("Datos no validos", true).Render(c.Request().Context(), c.Response().Writer)
		}

		db := config.DB
		user := dbmodels.User{}

		result := db.Where("username = ?", loginPayload.Username).First(&user)

		if result.Error != nil {
			c.Logger().Error(result.Error.Error())
			c.Response().WriteHeader(http.StatusNotFound)
			return layout.SimpleNotification("Usuario no encontrado", true).Render(c.Request().Context(), c.Response().Writer)
		}

		if user.Password != loginPayload.Password {
			c.Response().WriteHeader(http.StatusUnauthorized)
			return layout.SimpleNotification("Usuario o contraseña incorrecta", true).Render(c.Request().Context(), c.Response().Writer)
		}

		token, err := config.GetJwtToken(&networkmodels.JwtClaims{
			Sub:  user.ID,
			Role: int(user.RoleID),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		})

		if err != nil {
			c.Logger().Error(err)
			c.Response().WriteHeader(http.StatusInternalServerError)
			return layout.FailAlert("Error al generar token").Render(c.Request().Context(), c.Response().Writer)
		}

		c.SetCookie(&http.Cookie{
			Name:   "token",
			Value:  token,
			MaxAge: 86400,
		})

		if user.RoleID == 1 {
			return RenderTemplComp(pages.HxGetPage("/admin"))(c)
		} else if user.RoleID == 2 {
			return RenderTemplComp(pages.HxGetPage("/user"))(c)
		} else {
			return RenderTemplComp(pages.HxGetPage("/worker"))(c)
		}
	}
}

func RenderTemplComp(comp templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		return comp.Render(c.Request().Context(), c.Response().Writer)
	}
}

func verifyCreateUserPayload(payload networkmodels.CreateUserPayload) bool {
	return payload.Username == "" || payload.Password == "" || payload.RoleName == "" || payload.Name == "" || payload.LastName == ""
}

func generateAccountNumber(user dbmodels.User, accountType int) string {
	q := 0
	for _, r := range user.Username {
		q += int(r)
	}
	return strconv.Itoa(q)[0:3] + strconv.Itoa(accountType) + strconv.Itoa(int(user.ID))
}
