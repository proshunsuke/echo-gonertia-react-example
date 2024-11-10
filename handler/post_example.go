package handler

import (
	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

type PostExampleHandler struct {
	Base
}

func PostExample() *PostExampleHandler {
	return &PostExampleHandler{}
}

func (h *PostExampleHandler) Index(i *inertia.Inertia) echo.HandlerFunc {
	return func(c echo.Context) error {
		return i.Render(c.Response(), c.Request(), "PostExample/Index", inertia.Props{
			"text": "Post Example",
		})
	}
}

type Data struct {
	PostData string `json:"post_data" validate:"required,gte=1" validateMsg:"post_data is required"`
}

func (h *PostExampleHandler) Post(i *inertia.Inertia) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data Data
		ctx, err := h.ParseBody(c, &data)
		if err != nil {
			i.Redirect(c.Response(), c.Request().WithContext(ctx), "/post-example")
		}

		// Perform data store operation and other processing here

		i.Redirect(c.Response(), c.Request().WithContext(ctx), "/post-example")
		return nil
	}
}
