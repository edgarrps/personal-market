package main

import (
	"log"
	"os"

	"github.com/edgarrps/personal-market-backend/controller/routes"
	"github.com/edgarrps/personal-market-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
