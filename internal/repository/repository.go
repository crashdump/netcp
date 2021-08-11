package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/crashdump/netcp/internal/model"
)

var gdb *gorm.DB

// Open establishes connection to database and saves its handler into gdb *gorm.DB
func OpenGorm(connection string) {
	var err error
	gdb, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetGorm() *gorm.DB {
	return gdb
}

// AutoMigrate runs gorm auto migration
func AutoMigrate() {
	gdb.AutoMigrate(
		&model.User{},
	)
}