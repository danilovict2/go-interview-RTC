package database

import (
	"fmt"
	"os"

	"github.com/danilovict2/go-interview-RTC/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	migrate(db)

	return db, err
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Printf("Error migrating User model: %v\n", err)
	}

	if err := db.AutoMigrate(&models.Interview{}); err != nil {
		fmt.Printf("Error migrating Interview model: %v\n", err)
	}
}
