package migrations

import (
	"github.com/AdamHutchison/flux/database/models"
	"gorm.io/gorm"
)

func RegisterAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(new(models.User))
}

func RegisterStandardMigrations(db *gorm.DB) {
	
}