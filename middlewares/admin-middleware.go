package middlewares

import (
	"immersive/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := ExtractToken(c)
		if id > 0 && role == "admin" {
			return next(c)
		}
		return c.JSON(http.StatusForbidden, helpers.FailedResponse("Forbidden access not admin"))
	}
}
