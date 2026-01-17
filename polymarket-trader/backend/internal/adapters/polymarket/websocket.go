package polymarket

import (
	"log"

	"github.com/gorilla/websocket"
)

const WSURL = "wss://ws-subscriptions-clob.polymarket.com/ws/market"

type WSClient struct {
	Conn    *websocket.Conn
	MsgChan chan []byte
	Done    chan struct{}
}

func NewWSClient() *WSClient {
	return &WSClient{
		MsgChan: make(chan []byte, 100),
		Done:    make(chan struct{}),
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
		// In a real app, process specific event types here
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
