// Creates container
// Registers http kernal
package bootstrap

import (
	"fmt"
	"log"

	"github.com/AdamHutchison/flux/http"
	env "github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type FluxApp struct {
	kernal http.HttpKernal
}

func (f *FluxApp) Bootstrap() {
	bootstrapEnv()
	bootstrapConfig()
}

func (f *FluxApp) GetKernal() http.HttpKernal {
	return f.kernal
}

func bootstrapEnv() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(fmt.Errorf("Fatal error loading .env"))
	}
}

func bootstrapConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
