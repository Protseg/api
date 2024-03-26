package handler

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kazzkiq/protseg-api/database"
)

type SearchCustomerLine struct {
	ProductName string `json:"product_name"`
	ProductRef  string `json:"product_ref"`
}

// Greetings handle api status
func SearchByCustomer(c fiber.Ctx) error {
	searchCustomerId := strings.ToLower(c.Query("id"))
	db := database.DBConn

	if searchCustomerId == "" {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Busca não pode ser vazia, informe um cliente (id).",
		})
	}

	var searchResults []SearchCustomerLine
	var emptyResult = make([]SearchCustomerLine, 0)

	db.Raw("SELECT psimples.adReferencia AS product_ref, REPLACE(COALESCE(psimples.adDiscriminacao, pcomposto.adModelo), '\t', '') AS product_name FROM ppedidos pedidos LEFT JOIN ppedidosimples psimples ON pedidos.id = psimples.adIdPedido LEFT JOIN ppedidocomposto pcomposto ON pedidos.id = pcomposto.adIdPedido WHERE pedidos.adIdCliente = ? GROUP BY product_name", searchCustomerId).Scan(&searchResults)

	if searchResults == nil {
		searchResults = emptyResult
	}

	return c.JSON(fiber.Map{
		"success": true,
		"payload": searchResults,
	})
}
