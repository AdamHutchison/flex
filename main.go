package main

import (
	"github.com/AdamHutchison/flux/bootstrap"
)

func main() {
	app := bootstrap.FluxApp{}

	app.Bootstrap()

	kernal := app.GetKernal()

	kernal.HandleRequests()
}
