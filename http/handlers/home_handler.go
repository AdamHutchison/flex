package handlers

import (
	"net/http"

	"github.com/AdamHutchison/flux/http/transformers"
)

type HomeHandler struct {
	BaseHandler
}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := transformers.HomeTransformer{
		Message: "Welcome to your new flux app",
	}

	h.Respond(data, w, http.StatusOK)
}