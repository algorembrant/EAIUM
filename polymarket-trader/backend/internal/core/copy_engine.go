package core

import (
	"log"
	"sync"
	"time"

	"github.com/user/polymarket-trader/internal/adapters/polymarket"
	"github.com/user/polymarket-trader/internal/models"
)

type CopyEngine struct {
	Client          *polymarket.Client
	WS              *polymarket.WSClient
	Monitored       map[string]models.CopyConfig // Key: TraderAddress
	ActivePositions map[uint]*models.Position    // Key: Position ID (simulated for now)
	SignalChan      chan models.Position         // Ingest signals here
	mu              sync.RWMutex
	posIDCounter    uint
}

func NewCopyEngine(client *polymarket.Client, ws *polymarket.WSClient) *CopyEngine {
	return &CopyEngine{
		Client:          client,
		WS:              ws,
		Monitored:       make(map[string]models.CopyConfig),
		ActivePositions: make(map[uint]*models.Position),
		SignalChan:      make(chan models.Position, 100),
		posIDCounter:    1,
	}
}

func (e *CopyEngine) Start() {
	log.Println("Starting Copy Engine...")
	go e.processSignals()
	go e.processPriceUpdates()
}

func (e *CopyEngine) processPriceUpdates() {
	if e.WS == nil {
		return
	}
	for update := range e.WS.PriceUpdateChan {
		e.UpdateMarketPrice(update.MarketID, update.Price)
	}
}

func (e *CopyEngine) AddTrader(config models.CopyConfig) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Monitored[config.TraderAddress] = config
	log.Printf("Monitoring trader: %s", config.TraderAddress)
}

func (e *CopyEngine) GetActivePositions() []models.Position {
	e.mu.RLock()
	defer e.mu.RUnlock()
	var positions []models.Position
	for _, p := range e.ActivePositions {
		if p.IsOpen {
			positions = append(positions, *p)
		}
	}
	return positions
}

// UpdateMarketPrice triggers risk checks for all open positions in the given market
func (e *CopyEngine) UpdateMarketPrice(marketID string, newPrice float64) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, pos := range e.ActivePositions {
		if !pos.IsOpen || pos.MarketID != marketID {
			continue
		}

		pos.CurrentPrice = newPrice
		pos.Value = pos.Size * newPrice // Simplified value calc

		// Check Stop Loss
		if pos.StopLossPrice > 0 && newPrice <= pos.StopLossPrice {
			e.closePosition(pos, "Stop Loss Triggered")
		}

		// Check Take Profit
		if pos.TakeProfitPrice > 0 && newPrice >= pos.TakeProfitPrice {
			e.closePosition(pos, "Take Profit Triggered")
		}
	}
}

func (e *CopyEngine) closePosition(pos *models.Position, reason string) {
	log.Printf("Closing Position %d (%s): %s at price %.2f", pos.ID, pos.MarketID, reason, pos.CurrentPrice)
	// In real app: Send Sell Order

	// Mock close
	pos.IsOpen = false
	pos.UpdatedAt = time.Now()
}

func (e *CopyEngine) processSignals() {
	for signal := range e.SignalChan {
		e.executeCopy(signal)
	}
}

func (e *CopyEngine) executeCopy(signal models.Position) {
	e.mu.Lock()
	defer e.mu.Unlock()

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

	// Calculate Risk Params
	slPrice := 0.0
	tpPrice := 0.0
	if config.StopLossPct > 0 {
		slPrice = signal.EntryPrice * (1 - config.StopLossPct)
	}
	if config.TakeProfitPct > 0 {
		tpPrice = signal.EntryPrice * (1 + config.TakeProfitPct)
	}

	// Execution Logic
	order := polymarket.OrderRequest{
		TokenID: signal.MarketID,
		Price:   signal.EntryPrice,
		Side:    "BUY",
		Size:    size,
	}

	err := e.Client.PlaceOrder(order)
	if err != nil {
		log.Printf("Failed to copy trade: %v", err)
		return
	}

	// Track Position
	newPos := &models.Position{
		ID:              e.posIDCounter,
		TraderAddress:   signal.TraderAddress,
		MarketID:        signal.MarketID,
		EntryPrice:      signal.EntryPrice,
		CurrentPrice:    signal.EntryPrice,
		Size:            size,
		StopLossPrice:   slPrice,
		TakeProfitPrice: tpPrice,
		IsOpen:          true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	e.ActivePositions[e.posIDCounter] = newPos
	e.posIDCounter++

	log.Printf("Successfully copied trade for %s. tracked ID: %d", signal.TraderAddress, newPos.ID)
}
