package controllers

import (
	"github.com/evaizee/seat-arrangements/backend/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// SeatController handles HTTP requests related to seats
type SeatController struct {
	seatService services.SeatService
}

// NewSeatController creates a new SeatController
func NewSeatController(seatService services.SeatService) *SeatController {
	return &SeatController{
		seatService: seatService,
	}
}

// GetSeatMap handles GET /api/seats/map
func (c *SeatController) GetSeatMap(ctx *fiber.Ctx) error {
	// Get query parameters
	flightID := ctx.Query("flightId")
	passengerID := ctx.Query("passengerId")

	if flightID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Flight ID is required",
		})
	}

	// Get the seat map
	seatMap, err := c.seatService.GetSeatMap(flightID, passengerID)
	if err != nil {
		zap.L().Error("Failed to get seat map", zap.Error(err),
			zap.String("flight_id", flightID),
			zap.String("passenger_id", passengerID))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to get seat map",
		})
	}

	// Return the seat map
	return ctx.JSON(seatMap)
}
