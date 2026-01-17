package models

import (
	"time"
)

type Trader struct {
	Address     string  `json:"address" gorm:"primaryKey"`
	WinRate     float64 `json:"win_rate"`
	TotalPnL    float64 `json:"total_pnl"`
	TradeCount  int     `json:"trade_count"`
	LastActive  time.Time
	IsMonitored bool `json:"is_monitored"`
}

type Position struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	TraderAddress   string  `json:"trader_address" gorm:"index"`
	MarketID        string  `json:"market_id"`
	OutcomeIndex    int     `json:"outcome_index"`
	EntryPrice      float64 `json:"entry_price"`
	CurrentPrice    float64 `json:"current_price"`
	Size            float64 `json:"size"`
	Value           float64 `json:"value"`
	StopLossPrice   float64 `json:"stop_loss_price"`
	TakeProfitPrice float64 `json:"take_profit_price"`
	IsOpen          bool    `json:"is_open"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CopyConfig struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	TraderAddress string  `json:"trader_address" gorm:"uniqueIndex"`
	Enabled       bool    `json:"enabled"`
	FixedSize     float64 `json:"fixed_size"` // Amount to bet per trade
	MaxPosition   float64 `json:"max_position"`
	StopLossPct   float64 `json:"stop_loss_pct"`
	TakeProfitPct float64 `json:"take_profit_pct"`
}
