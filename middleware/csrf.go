package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomCSRF() echo.MiddlewareFunc {
	inertiaCSRFConfig := middleware.CSRFConfig{
		CookiePath:  "/",
		TokenLookup: "cookie:_csrf",
	}
	inertiaCSRF := middleware.CSRFWithConfig(inertiaCSRFConfig)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return inertiaCSRF(next)(c)
		}
	}
}
