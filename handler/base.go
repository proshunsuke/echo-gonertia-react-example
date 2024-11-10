package handler

import (
	"context"
	"echo-gonertia-react-example/lib"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
	"reflect"
)

type Base struct {
}

func (*Base) ParseBody(c echo.Context, m interface{}) (context.Context, error) {
	ctx := lib.NewEchoContext(c.Request().Context(), c)

	if err := c.Bind(m); err != nil {
		return ctx, err
	}
	if err := c.Validate(m); err != nil {
		validationErrors := make(inertia.ValidationErrors)
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				field, _ := reflect.TypeOf(m).Elem().FieldByName(fe.StructField())
				key := field.Tag.Get("json")
				if key == "" {
					key = fe.StructField()
				}
				validateMsg := field.Tag.Get("validateMsg")
				if validateMsg == "" {
					validateMsg = fe.Error()
				}
				validationErrors[key] = validateMsg
			}
		}
		ctx := inertia.SetValidationErrors(ctx, validationErrors)
		return ctx, err
	}
	return ctx, nil
}
