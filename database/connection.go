package database

import (
	"fmt"

	config "github.com/AdamHutchison/flux-config"
	"github.com/AdamHutchison/flux/database/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
}

func (c *Connection) RunMigrations() {
	db := DB()

	migrations.RegisterAutoMigrations(db)
	migrations.RegisterStandardMigrations(db)
}


func DB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: getDsn(),
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

func getDsn() string {
	user := config.Get("db.user")
	password := config.Get("db.password")
	host := config.Get("db.host")
	port := config.Get("db.port")
	database := config.Get("db.name")

	// username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
}
