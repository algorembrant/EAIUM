package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user/polymarket-trader/internal/core"
	"github.com/user/polymarket-trader/internal/models"
)

type Handler struct {
	Discovery  *core.TraderDiscoveryService
	CopyEngine *core.CopyEngine
}

func NewHandler(d *core.TraderDiscoveryService, c *core.CopyEngine) *Handler {
	return &Handler{
		Discovery:  d,
		CopyEngine: c,
	}
}

// GetTopTraders returns the list of discovered traders
func (h *Handler) GetTopTraders(c *fiber.Ctx) error {
	traders, err := h.Discovery.FetchTopTraders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(traders)
}

// StartCopying adds a trader to the copy engine
func (h *Handler) StartCopying(c *fiber.Ctx) error {
	var req models.CopyConfig
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	h.CopyEngine.AddTrader(req)
	return c.JSON(fiber.Map{"status": "monitoring", "trader": req.TraderAddress})
}

// GetPositions returns active positions
func (h *Handler) GetPositions(c *fiber.Ctx) error {
	positions := h.CopyEngine.GetActivePositions()
	return c.JSON(positions)
}
