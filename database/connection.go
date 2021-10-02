package database

import (
	"fmt"
	"log"
	"os"
	"time"

	config "github.com/AdamHutchison/flux-config"
	"github.com/AdamHutchison/flux/database/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	}), &gorm.Config{
		Logger: getDBLogger(),
	})

	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

func getDBLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)
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
