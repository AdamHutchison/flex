package validators

type HomeValidator struct {
	Name_string string `validate:"required,eq=40"`
}
