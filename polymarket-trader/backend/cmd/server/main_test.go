package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
	"github.com/user/polymarket-trader/internal/api"
	"github.com/user/polymarket-trader/internal/core"
	"github.com/user/polymarket-trader/internal/models"
)

// setupTestApp creates a fiber app with the same configuration as main, but with mocked/nil external clients
func setupTestApp() *fiber.App {
	// Initialize Core Services with nil clients for testing logic that doesn't strictly depend on them for these endpoints
	// Note: In a real scenario, we would interface these out and provide mocks.
	// For this basic smoke test, we rely on the fact that StartCopying doesn't immediately use the clients.
	copyEngine := core.NewCopyEngine(nil, nil)

	// We don't start the copy engine background routines to avoid nil pointer dereferences on clients

	discoveryService := core.NewTraderDiscoveryService()

	// Initialize Handler
	h := api.NewHandler(discoveryService, copyEngine)

	// Initialize Fiber app
	app := fiber.New()
	app.Use(cors.New())

	// Routes
	api.SetupRoutes(app, h)

	// Health Check (replicated from main)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"uptime": "running",
			"env":    "test",
		})
	})

	return app
}

func TestHealthCheck(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestStartCopying(t *testing.T) {
	app := setupTestApp()

	payload := models.CopyConfig{
		TraderAddress: "0xTestTrader",
		FixedSize:     10.0,
		Enabled:       true,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/api/copy", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetPositions(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("GET", "/api/positions", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
