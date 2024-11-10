package config

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonKey := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if jsonKey != "" && jsonKey != "-" {
			return jsonKey
		}
		return field.Name
	})
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
