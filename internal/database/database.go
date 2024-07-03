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

func ConnectDB() *gorm.DB {
	loadTheEnvfile()
	DB_URL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error making connection to the DATABASE")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Bank{})
	return db
}

func loadTheEnvfile() {
	// Assuming APP_ENV is set to either "debug" or "production"
	env := os.Getenv("APP_ENV")

	// Load the appropriate.env file based on the environment
	dir := os.Getenv("PROJECT_DIR")

	var err error
	// Construct the path to the .env.debug file
	if env == "debug" {
		envPath := filepath.Join(dir, ".env.debug")
		err = godotenv.Load(envPath)
	} else if env == "production" {
		envPath := filepath.Join(dir, ".env.production")
		err = godotenv.Load(envPath)
	} else {
		log.Fatal("Unsupported environment")
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
