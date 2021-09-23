// Creates container
// Registers http kernal
package bootstrap

import (
	config "github.com/AdamHutchison/flux-config"
	"github.com/AdamHutchison/flux/http"
)

type FluxApp struct {
	kernal http.HttpKernal
}

func (f *FluxApp) Bootstrap() {
	config.Load()
}

func (f *FluxApp) GetKernal() http.HttpKernal {
	return f.kernal
}

