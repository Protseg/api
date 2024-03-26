package handler

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kazzkiq/protseg-api/database"
)

type SearchReferenceLine struct {
	Reference string `json:"reference"`
	Unity     string `json:"unity"`
	Name      string `json:"name"`
	Ipi       string `json:"ipi"`
}

// Greetings handle api status
func SearchByReference(c fiber.Ctx) error {
	searchReference := strings.ToLower(c.Query("search"))
	db := database.DBConn

	if searchReference == "" {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Busca não pode ser vazia.",
		})
	}

	if len(searchReference) < 3 {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Texto da busca muito curto.",
		})
	}

	var searchResults []SearchReferenceLine
	var emptyResult = fiber.Map{}

	db.Raw("SELECT adReferencia as reference, adUnidade AS unity, adDiscriminacao AS name, adIpi AS ipi FROM ppedidosimples WHERE adReferencia = LOWER(?) LIMIT 1", searchReference).Scan(&searchResults)

	if searchResults == nil {
		return c.JSON(fiber.Map{
			"success": true,
			"payload": emptyResult,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"payload": searchResults[0],
	})
}
