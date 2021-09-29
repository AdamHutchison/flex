package bootstrap

import (
	log "github.com/sirupsen/logrus"
	config "github.com/AdamHutchison/flux-config"
	"github.com/AdamHutchison/flux/database"
	"github.com/AdamHutchison/flux/database/migrations"
)

type FluxApp struct {
	kernal HttpKernal
}

func (f *FluxApp) Bootstrap() {
	config.Load()
	f.configureLogger()
	f.kernal = NewKernal()
	f.migrateDB()
}

func (f *FluxApp) GetKernal() HttpKernal {
	return f.kernal
}

func (f *FluxApp) migrateDB() {
	connection := new(database.Connection)

	log.Info("Running Migrations ....")

	migrations.RegisterAutoMigrations(connection)
	connection.RunAutoMigrations()
	migrations.RegisterStandardMigrations(connection)

	log.Info("Migrations Complete")
}

func (f *FluxApp) configureLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
