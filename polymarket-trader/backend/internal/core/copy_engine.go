package core

import (
	"log"

	"github.com/user/polymarket-trader/internal/adapters/polymarket"
	"github.com/user/polymarket-trader/internal/models"
)

type CopyEngine struct {
	Client     *polymarket.Client
	WS         *polymarket.WSClient
	Monitored  map[string]models.CopyConfig // Key: TraderAddress
	SignalChan chan models.Position         // Ingest signals here
}

func NewCopyEngine(client *polymarket.Client, ws *polymarket.WSClient) *CopyEngine {
	return &CopyEngine{
		Client:     client,
		WS:         ws,
		Monitored:  make(map[string]models.CopyConfig),
		SignalChan: make(chan models.Position, 100),
	}
}

func (e *CopyEngine) Start() {
	log.Println("Starting Copy Engine...")
	go e.processSignals()
}

func (e *CopyEngine) AddTrader(config models.CopyConfig) {
	e.Monitored[config.TraderAddress] = config
	log.Printf("Monitoring trader: %s", config.TraderAddress)
}

func (e *CopyEngine) processSignals() {
	for signal := range e.SignalChan {
		e.executeCopy(signal)
	}
}

func (e *CopyEngine) executeCopy(signal models.Position) {
	config, exists := e.Monitored[signal.TraderAddress]
	if !exists || !config.Enabled {
		return
	}

	log.Printf("COPY TRIGGER: Trader %s entered market %s", signal.TraderAddress, signal.MarketID)

	// Calculate size
	size := config.FixedSize
	if size <= 0 {
		return
	}

	// Execution Logic
	// 1. Check current balance (skipped for brevity)
	// 2. Place order
	order := polymarket.OrderRequest{
		TokenID: signal.MarketID, // Simplified mapping
		Price:   signal.EntryPrice,
		Side:    "BUY", // Needs logic to determine side
		Size:    size,
	}

	err := e.Client.PlaceOrder(order)
	if err != nil {
		log.Printf("Failed to copy trade: %v", err)
	} else {
		log.Printf("Successfully copied trade for %s", signal.TraderAddress)
	}
}
