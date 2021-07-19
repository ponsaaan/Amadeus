package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// HealthCheck returns ok anyway.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "El psy kongroo")
}
