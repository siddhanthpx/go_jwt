package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var l log.Logger

func Register(c *fiber.Ctx) error {
	l.Println("Handling POST Request")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err

	}
	return c.JSON(data)
}
