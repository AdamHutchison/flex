package validators

import "github.com/go-playground/validator"

type ValidatorInterface interface {
	Validate() error
}

type BaseValidator struct {
}

func (v *BaseValidator) validate(data interface{}) error {
	validator := validator.New()

	return validator.Struct(data)
}