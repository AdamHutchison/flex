package database

import (
	"fmt"
	"reflect"

	config "github.com/AdamHutchison/flux-config"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	autoMigrations []interface{}
}

func (c *Connection) AddAutoMigration(model interface{}) {
	models := append(c.autoMigrations, model)
	c.autoMigrations = models
}

func (c *Connection) RunAutoMigrations() {
	db := DB()

	for _, model := range c.autoMigrations {
		log.Info("Migrating " + reflect.TypeOf(model).String())
		db.AutoMigrate(model)
	}
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
