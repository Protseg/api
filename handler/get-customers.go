package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kazzkiq/protseg-api/database"
)

type CustomerLine struct {
	Id           string `json:"id"`
	CustomerName string `json:"customer_name"`
}

// Greetings handle api status
func GetCustomers(c fiber.Ctx) error {
	db := database.DBConn

	var searchResults []CustomerLine
	var emptyResult = make([]CustomerLine, 0)

	db.Raw("SELECT id AS id, COALESCE(NULLIF(adRazaoSocial, ''), adNomeFantasia) AS customer_name FROM pclientes ORDER BY customer_name DESC").Scan(&searchResults)

	if searchResults == nil {
		searchResults = emptyResult
	}

	return c.JSON(fiber.Map{
		"success": true,
		"payload": searchResults,
	})
}
