# 220 Breakout Strategy Documentation

## Overview
**File:** `220breakout.mq5`
**Timeframe:** M3 (3 Minutes) - *Hardcoded in logic*
**Type:** Daily Range Breakout

The "220 Breakout" strategy identifies a specific time range each day, calculates the High (HHR) and Low (LLR) of that range, and looks for a breakout with a specific validation filter ("Clean Path"). It takes a maximum of one trade per day.

## Trading Logic

### 1. Range Definition
The strategy defines a "Daily Range" during which it records the Highest High and Lowest Low.
*   **Range End:** Controlled by `InpRangeEndHour` (Default: 00:00 / Midnight).
*   **Range Duration:** **Hardcoded to 2 hours**.
    *   *Note:* The input `InpRangeStartHour` exists in the settings but is **ignored** by the logic. The start time is always calculated as `Range End Time - 2 hours`.
*   **Calculation:** At the end of the range, the EA scans the M3 candles inside this window to find:
    *   **HHR (Highest High of Range)** & its time (`timeHHR`).
    *   **LLR (Lowest Low of Range)** & its time (`timeLLR`).

### 2. Entry Conditions
The EA checks for signals on every tick, but only acts on completed M3 candles.

#### Buy Signal:
1.  **Breakout:** A candle closes **above** the HHR.
2.  **Filter (Clean Path):** The path between the original High (`timeHHR`) and the Breakout Candle must be "clear".
    *   Use `IsPathClear`: Checks that no other candles between the `timeHHR` and the current Breakout Candle have touched or exceeded the `HHR`. This ensures the breakout is fresh and not a retest of previous choppy price action.

#### Sell Signal:
1.  **Breakout:** A candle closes **below** the LLR.
2.  **Filter (Clean Path):** The path between the original Low (`timeLLR`) and the Breakout Candle must be "clear".
    *   Use `IsPathClear`: Checks that no other candles between the `timeLLR` and the current Breakout Candle have touched or exceeded the `LLR`.

### 3. Execution Rules
*   **One Trade Per Day:** Uses the global flag `tradeTakenToday` to ensure only one entry is made per day.
*   **Trading Window:** Entries are taken after the range is complete.

## Risk Management (Exit Strategy)

The strategy uses tight stops and high reward targeting.

*   **Stop Loss (SL):** Placed at the opposite end of the **Breakout Candle**.
    *   **Buy:** SL = Low of the Breakout Candle.
    *   **Sell:** SL = High of the Breakout Candle.
*   **Take Profit (TP):** Calculated based on the Reward Ratio.
    *   `TP Distance = (Entry Price - SL Price) * InpRewardRatio`.
    *   Default Ratio is 10:1 (`InpRewardRatio = 10`).
*   **Position Sizing:**
    *   Calculated dynamically based on `InpRiskPercent` (Default: 1.0%).
    *   Lot size is determined so that if the SL is hit, the loss equals roughly 1% of the account balance.

## Parameters (Inputs)

| Group | Parameter | Default | Description |
| :--- | :--- | :--- | :--- |
| **Risk Management** | `InpRiskPercent` | `1.0` | Risk per trade as % of Balance. |
| | `InpRewardRatio` | `10` | Multiplier for TP (e.g., 10 means 10R). |
| **Time Settings** | `InpRangeStartHour`| `22` | *Actually Unused in calculation (Overridden by 2h duration).* |
| | `InpRangeEndHour` | `0` | Hour when the range calculation ends (0 = Midnight). |
| | `InpBrokerOffset` | `0` | Adjusts server time if needed. |
| **Visuals** | `InpDrawRange` | `true` | Draws the range box on the chart. |
| | `InpBoxColor` | `clrLavender` | Color of the range box. |
| **System** | `InpMagicNum` | `123456` | Magic number to identify EA orders. |

## Technical Notes
*   **Clean Path Algorithm:** The `IsPathClear` function iterates entirely through M3 bars between the Extreme Point (EP) and the Breakout Candle. This might be computationally intensive if the breakout happens very late in the day, but safeguards against "messy" breakouts.
*   **Hardcoded Timeframe:** Logic explicitly calls `CopyRates(_Symbol, PERIOD_M3, ...)`. Running the EA on an M15 or H1 chart will not change the logic; it will still look at M3 data internally.
