package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	t "github.com/AdamHutchison/flux/http/transformers"
	v "github.com/AdamHutchison/flux/http/validators"
	"github.com/go-playground/validator"
)

type BaseHandler struct {
	v.BaseValidator
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
		t.ErrorTransformer{Message: "The given data was invalid.", Errors: errorMessages})
}

func (bh *BaseHandler) Respond(data interface{}, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(t.BaseTransformer{Data: data})
}
