package utils

import (
	"auth-system/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbFile := os.Getenv("DB_FILE")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	DB.AutoMigrate(&models.User{})

}
