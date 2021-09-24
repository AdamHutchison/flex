package bootstrap

import (
	config "github.com/AdamHutchison/flux-config"
)

type FluxApp struct {
	kernal HttpKernal
}

func (f *FluxApp) Bootstrap() {
	config.Load()
	f.kernal = NewKernal()
}

func (f *FluxApp) GetKernal() HttpKernal {
	return f.kernal
}
