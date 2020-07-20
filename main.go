package main

import (
	"github.com/ardatosun/go-rest-api/database"
	"github.com/ardatosun/go-rest-api/routes"
	"github.com/gofiber/fiber"
	"log"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	routes.InitUserRoutes(app)

	err := app.Listen(3000)
	if err != nil {
		log.Panic("Failed to start Fiber app")
	}
}