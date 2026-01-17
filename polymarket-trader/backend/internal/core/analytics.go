package core

import (
	"github.com/user/polymarket-trader/internal/models"
)

// CalculateTraderPerformance computes WinRate and Total PnL from a history of trades
func CalculateTraderPerformance(trades []models.HistoricalTrade) (float64, float64, int) {
	if len(trades) == 0 {
		return 0, 0, 0
	}

	totalPnL := 0.0
	wins := 0
	count := 0

	// We assume 'trades' contains realized trades (sells/exits) that have a PnL value
	// Or we might need to group buys/sells to calculate realized PnL.
	// For simplicity, we assume the input is a list of "Closed Positions" or "Realized Trades" with PnL pre-calculated or calculable.

	for _, t := range trades {
		// Only count trades that are "exits" or have realized PnL
		if t.PnL != 0 {
			totalPnL += t.PnL
			if t.PnL > 0 {
				wins++
			}
			count++
		}
	}

	winRate := 0.0
	if count > 0 {
		winRate = float64(wins) / float64(count)
	}

	return winRate, totalPnL, count
}
