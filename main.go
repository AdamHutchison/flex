package main

import (
	"fmt"

	"github.com/AdamHutchison/flux/bootstrap"
	"github.com/spf13/viper"
)

func main() {
	// bootstraps the APP

	app := bootstrap.FluxApp{}

	app.Bootstrap()

	fmt.Println(viper.GetString("app.name"))

	// Resolves the kernal out the container

	// passes the request to the kernal

	// returns the response from the kernal
}
