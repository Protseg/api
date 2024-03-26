package main

import (
	"fmt"
	"os"

	"github.com/kazzkiq/protseg-api/database"
	"github.com/kazzkiq/protseg-api/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func getPort() string {
	port := os.Getenv("PORT")
	environment := os.Getenv("API_ENVIRONMENT")

	if port != "" {
		return ":" + port
	}

	if environment == "production" {
		return ":3001"
	} else if environment == "staging" {
		return ":3002"
	} else {
		return ":3003"
	}
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(database.ConnString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	initDatabase()

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

	port := getPort()
	err := app.Listen(port)
	panic(err)
}
