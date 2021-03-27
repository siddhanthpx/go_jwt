package main

import (
	"go-jwt/database"
	"go-jwt/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
