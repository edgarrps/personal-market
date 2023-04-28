package database

import (
	"log"
	"os"

	"github.com/edgarrps/personal-market-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database")
	} else {
		log.Println("Connect sucessfully")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
	)
}
