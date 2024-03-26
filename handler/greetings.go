package handler

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

// Greetings handle api status
func Greetings(c fiber.Ctx) error {
	environment := os.Getenv("API_ENVIRONMENT")

	return c.JSON(fiber.Map{
		"greetigs":    "Hello from Protseg API.",
		"environment": environment,
	})
}
