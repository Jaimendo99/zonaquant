package config

import (
	networkmodels "zonaquant/models/networkModels"
	"zonaquant/ui/authui"
	"zonaquant/ui/layout"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var JwtConfig = echojwt.Config{
	NewClaimsFunc: func(e echo.Context) jwt.Claims {
		return &networkmodels.JwtClaims{}
	},
	SigningKey:  []byte("secret"),
	TokenLookup: "cookie:token",
	ErrorHandler: func(e echo.Context, err error) error {
		println(err.Error())
		return authui.LoginComp().Render(e.Request().Context(), e.Response().Writer)
	},
}

// middleware to verify jwt token

func AuthorizationMiddleware(roleId int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*networkmodels.JwtClaims)
			if claims.Role > roleId {
				if c.Request().Header.Get("Hx-Request") == "" {
					return layout.MainLayout("No autorizado", layout.FailAlert("No autorizado")).Render(c.Request().Context(), c.Response().Writer)
				}
				return layout.FailAlert("No autorizado").Render(c.Request().Context(), c.Response().Writer)
			}
			return next(c)
		}
	}
}

func GetJwtToken(claim *networkmodels.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := "secret"
	println(secret)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
