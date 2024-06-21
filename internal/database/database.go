package database

import (
	"log"
	"os"
	"path/filepath"
	models "simple-bank/internal/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	// Assuming APP_ENV is set to either "debug" or "production"
	env := os.Getenv("APP_ENV")
	dir := os.Getenv("PROJECT_DIR")
	// Load the appropriate.env file based on the environment
	var err error
	if env == "debug" {
		// Construct the path to the .env.debug file
		envPath := filepath.Join(dir, ".env.debug")
		err = godotenv.Load(envPath)
	} else if env == "production" {
		// Construct the path to the .env.debug file
		envPath := filepath.Join(dir, ".env.production")
		err = godotenv.Load(envPath)
	} else {
		log.Fatal("Unsupported environment")
	}
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
