package routes

import (
	"github.com/evaizee/seat-arrangements/backend/di"
	// "github.com/evaizee/seat-arrangements/backend/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App, container *di.Container) {
	// API group
	api := app.Group("/api")

	// Health check endpoint
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok man",
		})
	})

	// Seat routes
	seat := api.Group("/seats")
	seat.Get("/map", container.SeatController.GetSeatMap)
}
