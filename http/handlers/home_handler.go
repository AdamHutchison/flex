package handlers

import (
	"net/http"

	"github.com/AdamHutchison/flux/http/transformers"
	v "github.com/AdamHutchison/flux/http/validators"
)

type HomeHandler struct {
	BaseHandler
}

func (h HomeHandler) Show(w http.ResponseWriter, r *http.Request) {
	err := h.Validate(new(v.HomeValidator), w, r)

	if err != nil {
		h.Error(err, w, http.StatusBadRequest)
		return
	}

	data := transformers.HomeTransformer{
		Message: "Welcome to your new flux app",
	}

	h.Respond(data, w, http.StatusOK)
}