package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/user/polymarket-trader/internal/adapters/polymarket"
	"github.com/user/polymarket-trader/internal/api"
	"github.com/user/polymarket-trader/internal/config"
	"github.com/user/polymarket-trader/internal/core"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	// Initialize Adapters
	pmClient := polymarket.NewClient(cfg)
	pmWS := polymarket.NewWSClient()

	// Start WS connection (in background)
	if err := pmWS.Connect(); err != nil {
		log.Printf("Warning: Failed to connect to Polymarket WS: %v", err)
	}

	// Initialize Core Services
	copyEngine := core.NewCopyEngine(pmClient, pmWS)
	copyEngine.Start()

	discoveryService := core.NewTraderDiscoveryService()

	// Initialize Handler
	h := api.NewHandler(discoveryService, copyEngine)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Polymarket Trader Engine",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	api.SetupRoutes(app, h)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"uptime": "running",
			"env":    cfg.ServerPort,
		})
	})

	// Start Server
	log.Printf("Starting server on port %s", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
