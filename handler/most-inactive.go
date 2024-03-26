package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/kazzkiq/protseg-api/database"
)

type MostInactiveLine struct {
	CustomerId    int32  `json:"id"`
	CustomerName  string `json:"customer_name"`
	LastOrderDate string `json:"last_order_date"`
	OrderId       int32  `json:"order_id"`
}

// Greetings handle api status
func GetMostInactive(c fiber.Ctx) error {
	numberOfDays, _ := strconv.Atoi(c.Query("days", "0"))
	providerId, _ := strconv.Atoi(c.Query("provider", "-1"))
	db := database.DBConn

	if numberOfDays == 0 {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Parametro (days) precisa ser maior que 0.",
		})
	}

	var results []MostInactiveLine
	var emptyResult = make([]MostInactiveLine, 0)

	if providerId == -1 {
		db.Raw("SELECT c.id AS customer_id, c.adRazaoSocial AS customer_name, (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) AS last_order_date, (SELECT id FROM ppedidos pp WHERE pp.adIdCliente = c.id ORDER BY pp.timestamp DESC LIMIT 1) AS order_id FROM pclientes c WHERE c.id NOT IN (SELECT c1.id FROM pclientes c1, ppedidos p WHERE p.adIdCliente = c1.id AND p.adData >= (date(DATE_SUB(SYSDATE(), INTERVAL ? DAY)))) AND (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) IS NOT NULL ORDER BY last_order_date DESC", numberOfDays).Scan(&results)
	} else {
		db.Raw("SELECT c.id AS customer_id, c.adRazaoSocial AS customer_name, (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) AS last_order_date, (SELECT id FROM ppedidos pp WHERE pp.adIdCliente = c.id ORDER BY pp.timestamp DESC LIMIT 1) AS order_id FROM pclientes c WHERE c.id NOT IN (SELECT c1.id FROM pclientes c1, ppedidos p WHERE p.adIdCliente = c1.id AND p.adData >= (date(DATE_SUB(SYSDATE(), INTERVAL ? DAY)))) AND (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) IS NOT NULL AND (SELECT COUNT(adIdFornecedor) FROM ppedidos pp WHERE pp.adIdCliente = c.id AND pp.adIdFornecedor != ?) = 0 ORDER BY last_order_date DESC", numberOfDays, providerId).Scan(&results)
	}

	if results == nil {
		results = emptyResult
	}

	return c.JSON(fiber.Map{
		"success": true,
		"payload": results,
	})
}
