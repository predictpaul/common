# PredictPaul Common Types

Common type definitions for Polymarket and Kalshi prediction market APIs.

## Installation

```bash
go get github.com/predictpaul/common
```

## Usage

### Unified API Response

```go
import "github.com/predictpaul/common"

// Parse API response with generics
var resp common.Response[polymarket.Order]
json.Unmarshal(body, &resp)

if resp.Code == common.CodeSuccess {
    order := resp.Data
    fmt.Printf("Order ID: %s\n", order.ID)
}

// Paginated response
var listResp common.Response[common.PageData[kalshi.OrderResponse]]
json.Unmarshal(body, &listResp)

for _, order := range listResp.Data.Data {
    fmt.Printf("Order: %s\n", order.OrderID)
}
```

### Polymarket Types

```go
import "github.com/predictpaul/common/polymarket"

// Market data from CLOB API
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

// API Response types
var order polymarket.Order
var balance polymarket.BalanceResponse
var positions []polymarket.Position
```

### Kalshi Types

```go
import "github.com/predictpaul/common/kalshi"

// Market data
var market kalshi.Market
fmt.Printf("Ticker: %s, Status: %s\n", market.Ticker, market.Status)

// API Response types
var orderResp kalshi.OrderResponse
var accountResp kalshi.AccountResponse
var flowResp kalshi.FlowResponse

// Check order status
if orderResp.Status == kalshi.APIOrderStatusFilled {
    fmt.Println("Order filled")
}

// Orderbook
var orderbook kalshi.Orderbook
yesLevels := orderbook.GetYesLevels()
for _, level := range yesLevels {
    fmt.Printf("Price: %d, Count: %d\n", level.Price, level.Count)
}
```

### Admin Types

```go
import "github.com/predictpaul/common/admin"

// Node configuration
var node admin.Node
if node.Status == admin.NodeStatusEnabled {
    fmt.Printf("Node %s is enabled\n", node.URL)
}

// Wallet configuration
var wallet admin.Wallet

// Market configuration
var marketConfig admin.MarketConfig
```

## Type Reference

### Common (github.com/predictpaul/common)

| Type | Description |
|------|-------------|
| `Response[T]` | Unified API response wrapper |
| `PageData[T]` | Paginated data wrapper |
| `CodeSuccess` | Success code (0) |
| `CodeFailed` | Failed code (101) |
| `CodeUnauthorized` | Unauthorized code (102) |

### Polymarket Platform Types (github.com/predictpaul/common/polymarket)

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

### Polymarket API Response Types

| Type | Description |
|------|-------------|
| `Order` | Order response data |
| `OrderListResponse` | Order list with pagination |
| `CancelBatchResponse` | Batch cancel result |
| `SyncResponse` | Sync result |
| `Account` | User account data |
| `TokenAccount` | Token position data |
| `WithdrawResponse` | Withdraw result |
| `SettleResult` | Single settle result |
| `BalanceResponse` | Balance summary |
| `Position` | Position with PnL info |

### Kalshi Platform Types (github.com/predictpaul/common/kalshi)

| Type | Description |
|------|-------------|
| `Side` | Position side (yes, no) |
| `Action` | Order action (buy, sell) |
| `OrderType` | Order type (limit, market) |
| `OrderStatus` | Order status (resting, canceled, executed, pending) |
| `MarketStatus` | Market status |
| `Market` | Market data |
| `Order` | Order data |
| `Position` | User position |
| `Balance` | Account balance |
| `Fill` | Trade fill |
| `Settlement` | Market settlement |
| `Orderbook` | Order book data |
| `CreateOrderParams` | Order creation parameters |

### Kalshi API Response Types

| Type | Description |
|------|-------------|
| `OrderResponse` | Order response data |
| `AccountResponse` | Account response data |
| `FlowResponse` | Fund flow record |
| `MarketResponse` | Market response data |
| `MarketListResponse` | Market list with cursor |
| `OrderbookResponse` | Orderbook data |
| `SettleResponse` | Market settle result |

### Kalshi Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `FlowTypeRecharge` | 1 | Deposit |
| `FlowTypeWithdraw` | 2 | Withdraw |
| `FlowTypeBuy` | 3 | Buy |
| `FlowTypeSell` | 4 | Sell |
| `FlowTypeSettle` | 5 | Settlement |
| `FlowTypeFee` | 6 | Fee |
| `FlowTypeFreeze` | 7 | Freeze |
| `FlowTypeUnfreeze` | 8 | Unfreeze |
| `APIOrderStatusPending` | 0 | Pending |
| `APIOrderStatusResting` | 1 | Resting |
| `APIOrderStatusFilled` | 2 | Filled |
| `APIOrderStatusCanceled` | 3 | Canceled |
| `APIOrderStatusPartiallyFilled` | 4 | Partially filled |
| `APIOrderStatusSettled` | 5 | Settled |

### Admin Types (github.com/predictpaul/common/admin)

| Type | Description |
|------|-------------|
| `Node` | Blockchain node configuration |
| `Wallet` | Wallet configuration |
| `MarketConfig` | Market account configuration |

## License

MIT License
