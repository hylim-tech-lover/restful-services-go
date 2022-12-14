package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hylim-tech-lover/restful-services-go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFact)

	app.Post("/fact", handlers.CreateFact)
}
