package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kazzkiq/protseg-api/interfaces"
)

type Logged struct {
	IsLogged bool `json:"logged"`
}

// Greetings handle api status
func IsLogged(c fiber.Ctx) error {
	sid := c.Query("session_id")

	if sid == "" {
		return c.Status(400).JSON(interfaces.ErrorResponse{
			Success: false,
			Message: "Parametro 'session_id' é necessário.",
		})
	}

	return c.JSON(interfaces.APIResponse{
		Success: true,
		Payload: false,
	})
}
