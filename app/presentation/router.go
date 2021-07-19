package presentation

import (
	echo "github.com/labstack/echo/v4"
	"github.com/ponsaaan/Amadeus.git/app/presentation/handler"
	"github.com/ponsaaan/Amadeus.git/app/presentation/middleware"
)

// Router returns http router.
func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSHeaders())
	e.Use(middleware.CORSValidator())

	// `/` endpoint is used to healthcheck.
	e.GET("/", handler.HealthCheck)

	return e
}
