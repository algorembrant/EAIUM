package polymarket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/user/polymarket-trader/internal/config"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Config     *config.Config
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		BaseURL: cfg.PolymarketAPI,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Config: cfg,
	}
}

// Basic struct for placing an order (simplified for Clob)
type OrderRequest struct {
	TokenID string  `json:"token_id"`
	Price   float64 `json:"price"`
	Side    string  `json:"side"` // "BUY" or "SELL"
	Size    float64 `json:"size"`
	FeeRate int     `json:"fee_rate_bps"`
}

// PlaceOrder sends an order to the CLOB
func (c *Client) PlaceOrder(req OrderRequest) error {
	// TODO: Implement signing logic using EIP-712
	// This is a placeholder for the structure

	url := fmt.Sprintf("%s/order", c.BaseURL)
	jsonData, _ := json.Marshal(req)

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	// Add Auth headers (Polymarket-ApiKey, Signature, etc.)

	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("order failed with status: %d", resp.StatusCode)
	}

	return nil
}
