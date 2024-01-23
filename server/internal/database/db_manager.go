package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func SetupDatabase() *gorm.DB {
	// Set up the test database
	db, err = gorm.Open(
		sqlite.Open("./quotation.db?cache=shared&loc=UTC"),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	MigrateDatabase(db)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func CloseDatabase(db *gorm.DB) {
	instance, _ := db.DB()
	err := instance.Close()
	if err != nil {
		log.Fatal("Failed to close database connection:", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
