// Creates container
// Registers http kernal
package bootstrap

import (
	"github.com/AdamHutchison/flux/http"
)

type FluxApp struct {
	kernal http.HttpKernal
}

func (f *FluxApp) Bootstrap() {
	config := Config{}

	config.LoadConfig()
}

func (f *FluxApp) GetKernal() http.HttpKernal {
	return f.kernal
}

