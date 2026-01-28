# PredictPaul Common Types

Common type definitions for Polymarket and Kalshi prediction market APIs.

## Installation

```bash
go get github.com/predictpaul/common
```

## Usage

### Polymarket Types

```go
import "github.com/predictpaul/common/polymarket"

// Market data
var market polymarket.ClobMarket

// Check market status
if market.IsClosed() {
    fmt.Println("Market is closed")
}

if market.IsSettled() {
    result := market.GetResult() // "Yes" or "No"
    fmt.Printf("Winner: %s\n", result)
}

// Get unified status: "open", "closed", or "settled"
status := market.GetUnifiedStatus()

// Trade data
var trade polymarket.Trade
fmt.Printf("Trade ID: %s, Side: %s, Size: %s\n", trade.ID, trade.Side, trade.Size)
```

### Kalshi Types

```go
import "github.com/predictpaul/common/kalshi"

// Market data
var market kalshi.Market
fmt.Printf("Ticker: %s, Status: %s\n", market.Ticker, market.Status)

// Order data
var order kalshi.Order
fmt.Printf("Order ID: %s, Side: %s, Action: %s\n", order.OrderID, order.Side, order.Action)

// Position data
var position kalshi.Position
fmt.Printf("Ticker: %s, Position: %d\n", position.Ticker, position.Position)

// Orderbook
var orderbook kalshi.Orderbook
yesLevels := orderbook.GetYesLevels()
noLevels := orderbook.GetNoLevels()
for _, level := range yesLevels {
    fmt.Printf("Price: %d, Count: %d\n", level.Price, level.Count)
}
```

## Type Reference

### Polymarket

| Type | Description |
|------|-------------|
| `OrderStatus` | Order status enum (PENDING, MATCHED, UNMATCHED, LIVE, DELAY, CANCELED) |
| `PolymarketEvent` | Market event data |
| `PolymarketMarket` | Market data from Gamma API |
| `ClobMarket` | Market data from CLOB API |
| `ClobToken` | Token info (token_id, outcome, price, winner) |
| `Trade` | Trade record |
| `TradeParams` | Trade query parameters |
| `TradesResponse` | Paginated trade list response |

### Kalshi

| Type | Description |
|------|-------------|
| `Side` | Position side (yes, no) |
| `Action` | Order action (buy, sell) |
| `OrderType` | Order type (limit, market) |
| `OrderStatus` | Order status (resting, canceled, executed, pending) |
| `MarketStatus` | Market status (initialized, inactive, active, closed, determined, disputed, amended, finalized) |
| `Market` | Market data |
| `Order` | Order data |
| `Position` | User position |
| `Balance` | Account balance |
| `Fill` | Trade fill |
| `Settlement` | Market settlement |
| `Orderbook` | Order book data |
| `CreateOrderParams` | Order creation parameters |

## License

MIT License
