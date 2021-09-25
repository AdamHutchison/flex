package validators

type HomeValidator struct {
	BaseValidator
	Name_string string `validate:"required,eq=40"`
}

func (v *HomeValidator) Validate() error {
	return v.BaseValidator.validate(v)
}