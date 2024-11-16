package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
	"net/http"
)

type ErrorHandler struct {
	inertia *inertia.Inertia
}

func Error(i *inertia.Inertia) *ErrorHandler {
	return &ErrorHandler{
		inertia: i,
	}
}

func (h *ErrorHandler) Error(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}
	c.Logger().Error(err)

	// Gonertia enforces a 200 status, so constructing response header manually.
	if inertia.IsInertiaRequest(c.Request()) {
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("X-Inertia", "true")
	} else {
		c.Response().Header().Set("Content-Type", "text/html")
	}
	c.Response().WriteHeader(code)

	err = h.inertia.Render(c.Response(), c.Request(), "Error/index", inertia.Props{
		"status": code,
	})

	if err != nil {
		c.Logger().Error(err)
	}
}
