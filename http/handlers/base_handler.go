package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	t "github.com/AdamHutchison/flux/http/transformers"
	v "github.com/AdamHutchison/flux/http/validators"
	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
)

type validateable interface {
	getValidator() v.ValidatorInterface
}

type BaseHandler struct {
}

func (bh *BaseHandler) validate(h validateable, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return bh.validateQuery(h, w, r)
	}

	return bh.validateBody(h, w, r)
}

func (bh *BaseHandler) validateBody(h validateable, w http.ResponseWriter, r *http.Request) error {
	validator := h.getValidator()

	err := json.NewDecoder(r.Body).Decode(validator)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return validator.Validate()
}

func (bh *BaseHandler) validateQuery(h validateable, w http.ResponseWriter, r *http.Request) error {
	validator := h.getValidator()

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := decoder.Decode(validator, r.URL.Query())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}

	return validator.Validate()
}

func (bh *BaseHandler) Error(err error, w http.ResponseWriter, statusCode int) {
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := make(map[string]string)

	for _, v := range validationErrors {
		field := strings.ToLower(v.Field())
		errorMessages[field] = "field " + field + " validation failed on the following rules: " + v.ActualTag() 
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(
		t.ErrorTransformer{ Message: "The given data was invalid.", Errors: errorMessages, })
}

func (bh *BaseHandler) Respond(data interface{}, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(t.BaseTransformer{Data: data})
}
