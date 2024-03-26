package handler

import (
	"github.com/gofiber/fiber/v3"
)

// Ping handle api status
func Ping(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"o": 1,
	})
}
