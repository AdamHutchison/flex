package bootstrap

import (
	log "github.com/sirupsen/logrus"
	config "github.com/AdamHutchison/flux-config"
	"github.com/AdamHutchison/flux/database"
)

type FluxApp struct {
	kernal HttpKernal
}

func (f *FluxApp) Bootstrap() {
	config.Load()
	f.configureLogger()
	f.migrateDB()
	f.kernal = NewKernal()
}

func (f *FluxApp) GetKernal() HttpKernal {
	return f.kernal
}

func (f *FluxApp) migrateDB() {
	connection := new(database.Connection)
	connection.RunMigrations()
}

func (f *FluxApp) configureLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
