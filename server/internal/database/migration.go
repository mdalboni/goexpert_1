package database

import (
	"goexpert_server_1/internal/models"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Quotation{})
}
