package handler

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kazzkiq/protseg-api/database"
)

type SearchProductLine struct {
	ID             int     `json:"id"`
	OrderId        string  `json:"order_id"`
	OrderTimestamp int     `json:"order_timestamp"`
	CustomerName   string  `json:"customer_name"`
	OrderValue     float64 `json:"order_value"`
	OrderType      int     `json:"order_type"`
}

// Greetings handle api status
func SearchByProduct(c fiber.Ctx) error {
	searchText := strings.ToLower(c.Query("search"))
	searchTextLike := "%" + searchText + "%"
	db := database.DBConn

	if searchText == "" {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Busca não pode ser vazia.",
		})
	}

	if len(searchText) < 3 {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Texto da busca muito curto.",
		})
	}

	var searchResults []SearchProductLine
	var emptyResult = make([]SearchProductLine, 0)

	db.Raw("SELECT pedidos.id AS id, pedidos.timestamp as order_timestamp, pedidos.idpedido AS order_id, COALESCE(NULLIF(clientes.adRazaoSocial, ''), clientes.adNomeFantasia) AS customer_name, pedidos.adValorTotal AS order_value, pedidos.adTipoPedido AS order_type FROM ppedidos pedidos LEFT JOIN ppedidosimples psimples ON pedidos.id = psimples.adIdPedido LEFT JOIN ppedidocomposto pcomposto ON pedidos.id = pcomposto.adIdPedido INNER JOIN pclientes clientes ON pedidos.adIdCliente = clientes.id WHERE LOWER(psimples.adReferencia) = ? OR LOWER(psimples.adDiscriminacao) LIKE ? OR LOWER(pcomposto.adModelo) LIKE ? GROUP BY pedidos.id ORDER BY pedidos.id DESC", searchText, searchTextLike, searchTextLike).Scan(&searchResults)

	if searchResults == nil {
		searchResults = emptyResult
	}

	return c.JSON(fiber.Map{
		"success": true,
		"payload": searchResults,
	})
}
