package polymarket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

const WSURL = "wss://ws-subscriptions-clob.polymarket.com/ws/market"

type WSClient struct {
	Conn            *websocket.Conn
	MsgChan         chan []byte
	PriceUpdateChan chan PriceUpdate
	Done            chan struct{}
}

type PriceUpdate struct {
	MarketID string
	Price    float64
}

func NewWSClient() *WSClient {
	return &WSClient{
		MsgChan:         make(chan []byte, 100),
		PriceUpdateChan: make(chan PriceUpdate, 100),
		Done:            make(chan struct{}),
	}
}

func (w *WSClient) Connect() error {
	c, _, err := websocket.DefaultDialer.Dial(WSURL, nil)
	if err != nil {
		return err
	}
	w.Conn = c

	// Start reading loop
	go w.readLoop()
	return nil
}

func (w *WSClient) readLoop() {
	defer close(w.Done)
	for {
		_, message, err := w.Conn.ReadMessage()
		if err != nil {
			log.Println("ws read error:", err)
			return
		}

		// Parse message (Simplified for demo)
		// We optimistically try to parse as a price update
		// In a real app, we'd check event type first
		var update struct {
			Type   string  `json:"event_type"`
			Market string  `json:"asset_id"`
			Price  float64 `json:"price"` // Assuming normalized price
			// CLOB messages are more complex (Orders, Trades, etc.)
			// We'll treat this as a generic update container for now
		}

		if err := json.Unmarshal(message, &update); err == nil {
			if update.Market != "" && update.Price > 0 {
				w.PriceUpdateChan <- PriceUpdate{
					MarketID: update.Market,
					Price:    update.Price,
				}
			}
		}

		w.MsgChan <- message
	}
}

type SubscriptionMsg struct {
	Type    string   `json:"type"`
	Channel string   `json:"channel"`
	Assets  []string `json:"assets_ids,omitempty"`
}

func (w *WSClient) Subscribe(assetIDs []string) error {
	msg := SubscriptionMsg{
		Type:    "Market",       // Example type
		Channel: "price_change", // Example channel
		Assets:  assetIDs,
	}
	return w.Conn.WriteJSON(msg)
}
