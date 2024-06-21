package database

import (
	"log"
	"os"
	models "simple-bank/internal/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	// Assuming APP_ENV is set to either "debug" or "production"
	env := os.Getenv("APP_ENV")
	// Load the appropriate.env file based on the environment
	var err error
	if env == "debug" {
		err = godotenv.Load(".env.debug")
	} else if env == "production" {
		err = godotenv.Load(".env.production")
	} else {
		log.Fatal("Unsupported environment")
	}
	// err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_URL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error making connection to the DATABASE")
	}
	db.AutoMigrate(&models.User{})
	return db
}
