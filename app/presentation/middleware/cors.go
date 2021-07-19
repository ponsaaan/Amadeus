package middleware

import (
	"net/http"
	"regexp"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var allowedOrigins = []string{
	"http://localhost:8080",
	"http://localhost:8081",
}

var allowedPaths = []string{
	"/",
}

// CORSHeaders returns a CORS middleware with config.
func CORSHeaders() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc:  allowOrigin,
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
			echo.HeaderXCSRFToken,
			echo.HeaderAuthorization,
			echo.HeaderOrigin,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
		MaxAge: 86400,
	})
}

func allowOrigin(origin string) (bool, error) {
	for _, o := range allowedOrigins {
		r := regexp.MustCompile(o)
		if r.MatchString(origin) {
			return true, nil
		}
	}
	return false, nil
}

// CORSValidator validates a header origin.
func CORSValidator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			for _, p := range allowedPaths {
				if path == p {
					return next(c)
				}
			}
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			for _, o := range allowedOrigins {
				r := regexp.MustCompile(o)
				if r.MatchString(origin) {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, map[string]string{"status": "forbidden"})
		}
	}
}
