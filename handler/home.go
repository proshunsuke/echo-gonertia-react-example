package handler

import (
	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

type HomeHandler struct {
	Base
}

func Home() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Index(i *inertia.Inertia) echo.HandlerFunc {
	return func(c echo.Context) error {
		return i.Render(c.Response(), c.Request(), "index", inertia.Props{
			"text": "Hello, World!",
		})
	}
}
