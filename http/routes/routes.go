package routes

import (
	h "github.com/AdamHutchison/flux/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(mux *mux.Router) {
	mux.Handle("/", h.HomeHandler{}).Methods("GET").Name("home")
}