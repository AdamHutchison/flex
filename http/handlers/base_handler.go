package handlers

import (
	"encoding/json"
	"net/http"

	t "github.com/AdamHutchison/flux/http/transformers"
)

type BaseHandler struct {
}

func (h *BaseHandler) Respond(data interface{}, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(t.BaseTransformer{Data: data})
}
