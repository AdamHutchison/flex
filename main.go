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

	var config string
	viper.UnmarshalKey("app.name", &config, viper.DecodeHook(bootstrap.Hook))

	fmt.Println()

	// Resolves the kernal out the container

	// passes the request to the kernal

	// returns the response from the kernal
}
