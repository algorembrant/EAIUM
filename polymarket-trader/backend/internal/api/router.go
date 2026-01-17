package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *Handler) {
	api := app.Group("/api")

	api.Get("/traders", h.GetTopTraders)
	api.Post("/copy", h.StartCopying)
	api.Get("/positions", h.GetPositions)
}
