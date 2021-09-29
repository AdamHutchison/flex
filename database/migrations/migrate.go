package migrations

import (
	"github.com/AdamHutchison/flux/database"
	"github.com/AdamHutchison/flux/database/models"
)

func RegisterAutoMigrations(db *database.Connection) {
	db.AddAutoMigration(new(models.User))
}

func RegisterStandardMigrations(db *database.Connection) {
	
}