package middlewares

import (
	"strings"

	"github.com/iryoda/price-guru/app/providers"
	"github.com/labstack/echo/v4"
)

type AuthContext struct {
	echo.Context
	User string
}

type AuthMiddleware struct {
	TokenProvider providers.TokenProvider
}

func (m AuthMiddleware) WithJWTToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")

		token := strings.Split(header, "Bearer ")

		if token == nil || len(token) < 2 {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		claims, err := m.TokenProvider.GetClaims(token[1])

		if err != nil {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		c.Set("User", claims.UserId)

		return next(c)
	}
}
