package core

import (
	"log"
	"sort"
	"time"

	"github.com/user/polymarket-trader/internal/models"
)

type TraderDiscoveryService struct {
	// In a real app, this would use a GraphClient or DB connection
	topTraders []models.Trader
}

func NewTraderDiscoveryService() *TraderDiscoveryService {
	return &TraderDiscoveryService{
		topTraders: []models.Trader{},
	}
}

// FetchTopTraders mock implementation
func (s *TraderDiscoveryService) FetchTopTraders() ([]models.Trader, error) {
	log.Println("Fetching top traders from 'source'...")

	// Mock Data
	mockTraders := []models.Trader{
		{
			Address:     "0x9f8...7b2",
			WinRate:     0.78,
			TotalPnL:    12500.50,
			TradeCount:  142,
			LastActive:  time.Now().Add(-10 * time.Minute),
			IsMonitored: false,
		},
		{
			Address:     "0xa12...8c9",
			WinRate:     0.65,
			TotalPnL:    8900.20,
			TradeCount:  98,
			LastActive:  time.Now().Add(-2 * time.Hour),
			IsMonitored: true,
		},
		{
			Address:     "0xb34...1d4",
			WinRate:     0.92,
			TotalPnL:    45000.00,
			TradeCount:  310,
			LastActive:  time.Now().Add(-5 * time.Minute),
			IsMonitored: false,
		},
	}

	// Sort by PnL
	sort.Slice(mockTraders, func(i, j int) bool {
		return mockTraders[i].TotalPnL > mockTraders[j].TotalPnL
	})

	s.topTraders = mockTraders
	return mockTraders, nil
}
