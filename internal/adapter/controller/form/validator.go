package form

import (
	"gopkg.in/go-playground/validator.v9"
)

func Validate(form interface{}) error {
	validate := validator.New()

	err := validate.Struct(form)
	return err
}
