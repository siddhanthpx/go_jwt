package main

import (
	"go-jwt/database"
	"go-jwt/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connection to PSQL database
	err := database.ConnectDB()
	CheckError(err)
	log.Println("Successfully connected to database.")

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
