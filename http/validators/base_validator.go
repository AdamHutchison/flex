package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
)

type Validateable interface {
	GetValidator() interface{}
	http.Handler
}

type BaseValidator struct {
}

func (v *BaseValidator) Validate(h Validateable, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return v.validateQuery(h, w, r)
	}

	return v.validateBody(h, w, r)
}

func (v *BaseValidator) validateBody(h Validateable, w http.ResponseWriter, r *http.Request) error {
	validator := h.GetValidator()

	err := json.NewDecoder(r.Body).Decode(validator)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return validateStruct(validator)
}

func (v *BaseValidator) validateQuery(h Validateable, w http.ResponseWriter, r *http.Request) error {
	validator := h.GetValidator()

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	err := decoder.Decode(validator, r.URL.Query())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err)
	}

	return validateStruct(validator)
}

func validateStruct(data interface{}) error {
	validator := validator.New()

	return validator.Struct(data)
}
