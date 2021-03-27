package main

import (
	"go-jwt/database"
	"go-jwt/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Connection to PSQL database
	err := database.ConnectDB()
	checkError(err)
	log.Println("Successfully connected to database.")

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
