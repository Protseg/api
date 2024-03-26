package router

import (
	"github.com/kazzkiq/protseg-api/handler"

	"github.com/gofiber/fiber/v3"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Greetings)
	app.Get("/ping", handler.Ping)
	app.Get("/customers", handler.GetCustomers)
	app.Get("/search/by-customer", handler.SearchByCustomer)
	app.Get("/search/by-product", handler.SearchByProduct)
	app.Get("/search/by-reference", handler.SearchByReference)
	app.Get("/customers/most-inactive", handler.GetMostInactive)
}
