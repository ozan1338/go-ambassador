package main

import (
	"go-ambassador/src/database"
	"go-ambassador/src/events"
	"go-ambassador/src/routes"
	"go-ambassador/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()
	services.Setup()
	events.SetupProducer()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
