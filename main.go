package main

import (
	"github.com/AdityaAWP/go-crud/database"
	"github.com/AdityaAWP/go-crud/database/migration"
	"github.com/AdityaAWP/go-crud/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	routes.RouteInit(app)
	app.Listen(":3000")
}
