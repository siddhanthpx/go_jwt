package routes

import (
	"go-jwt/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/Register", controllers.Register)

}
