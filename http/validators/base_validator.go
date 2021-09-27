package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
)

type BaseValidator struct {
}

func (v *BaseValidator) Validate(validator interface{}, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return v.validateQuery(validator, w, r)
	}

	return v.validateBody(validator, w, r)
}

func (v *BaseValidator) validateBody(validator interface{}, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(validator)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return validateStruct(validator)
}

func (v *BaseValidator) validateQuery(validator interface{}, w http.ResponseWriter, r *http.Request) error {
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
