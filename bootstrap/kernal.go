package bootstrap

import (
	"log"
	"net/http"

	config "github.com/AdamHutchison/flux-config"
	"github.com/AdamHutchison/flux/http/middleware"
	"github.com/AdamHutchison/flux/http/routes"
	"github.com/gorilla/mux"
)

type HttpKernal struct {
	router *mux.Router
}

func (k *HttpKernal) HandleRequests() {
	log.Fatal(http.ListenAndServe(":"+config.Get("app.port"), k.router))
}

func (k *HttpKernal) GetRouter() *mux.Router {
	return k.router
}

func NewKernal() HttpKernal {
	kernal := HttpKernal{
		router: mux.NewRouter(),
	}

	routes.RegisterRoutes(kernal.router)
	middleware.RegisterGlobalMiddleware(kernal.router)

	return kernal
}
