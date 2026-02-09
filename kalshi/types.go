// Package kalshi provides common types for the Kalshi prediction market API.
// This package can be imported by third-party applications via:
//
//	go get github.com/predictpaul/common/kalshi
package kalshi

import "time"

// Side represents the side of a position (yes/no).
type Side string

const (
	SideYes Side = "yes"
	SideNo  Side = "no"
)

// Action represents the action of an order (buy/sell).
type Action string

const (
	ActionBuy  Action = "buy"
	ActionSell Action = "sell"
)

// OrderType represents the type of an order.
type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

// OrderStatus represents the status of an order.
type OrderStatus string

const (
	OrderStatusResting  OrderStatus = "resting"
	OrderStatusCanceled OrderStatus = "canceled"
	OrderStatusExecuted OrderStatus = "executed"
	OrderStatusPending  OrderStatus = "pending"
)

// TimeInForce represents the time-in-force policy for an order.
type TimeInForce string

const (
	TIFGoodTilCanceled   TimeInForce = "good_till_canceled"
	TIFFillOrKill        TimeInForce = "fill_or_kill"
	TIFImmediateOrCancel TimeInForce = "immediate_or_cancel"
)

// MarketStatus represents the status of a market.
// Available: initialized, inactive, active, closed, determined, disputed, amended, finalized
type MarketStatus string

const (
	MarketStatusInitialized MarketStatus = "initialized"
	MarketStatusInactive    MarketStatus = "inactive"
	MarketStatusActive      MarketStatus = "active"
	MarketStatusClosed      MarketStatus = "closed"
	MarketStatusDetermined  MarketStatus = "determined"
	MarketStatusDisputed    MarketStatus = "disputed"
	MarketStatusAmended     MarketStatus = "amended"
	MarketStatusFinalized   MarketStatus = "finalized"

	// Legacy aliases
	MarketStatusOpen    MarketStatus = "active"
	MarketStatusSettled MarketStatus = "finalized"
)

// Market represents a Kalshi market.
type Market struct {
	Ticker         string       `json:"ticker"`
	EventTicker    string       `json:"event_ticker"`
	Title          string       `json:"title"`
	Subtitle       string       `json:"subtitle"`
	Status         MarketStatus `json:"status"`
	Result         string       `json:"result"` // "yes", "no", or ""
	YesBid         int          `json:"yes_bid"`
	YesAsk         int          `json:"yes_ask"`
	NoBid          int          `json:"no_bid"`
	NoAsk          int          `json:"no_ask"`
	LastPrice      int          `json:"last_price"`
	Volume         int64        `json:"volume"`
	Volume24H      int64        `json:"volume_24h"`
	OpenInterest   int64        `json:"open_interest"`
	OpenTime       time.Time    `json:"open_time"`
	CloseTime      time.Time    `json:"close_time"`
	ExpirationTime time.Time    `json:"expiration_time"`
	TickSize       int          `json:"tick_size"`
	RiskLimitCents int64        `json:"risk_limit_cents"`
}

// IsSettled returns whether the market has been settled.
func (m *Market) IsSettled() bool {
	return m.Status == MarketStatusFinalized
}

// IsClosed returns whether the market is closed.
func (m *Market) IsClosed() bool {
	return m.Status == MarketStatusClosed ||
		m.Status == MarketStatusDetermined ||
		m.Status == MarketStatusFinalized
}

// GetUnifiedStatus returns a unified status string.
// Kalshi status mapping:
//   - active -> open
//   - finalized -> settled
//   - closed/determined -> closed
//   - others -> the original status
func (m *Market) GetUnifiedStatus() string {
	switch m.Status {
	case MarketStatusActive:
		return "open"
	case MarketStatusFinalized:
		return "settled"
	case MarketStatusClosed, MarketStatusDetermined:
		return "closed"
	default:
		return string(m.Status)
	}
}

// IsYesWinner returns whether the market result is "yes".
func (m *Market) IsYesWinner() bool {
	return m.Result == "yes"
}

// Order represents a Kalshi order.
type Order struct {
	OrderID              string      `json:"order_id"`
	UserID               string      `json:"user_id"`
	ClientOrderID        string      `json:"client_order_id"`
	Ticker               string      `json:"ticker"`
	Status               OrderStatus `json:"status"`
	Side                 Side        `json:"side"`
	Action               Action      `json:"action"`
	Type                 OrderType   `json:"type"`
	YesPrice             int         `json:"yes_price"`
	NoPrice              int         `json:"no_price"`
	YesPriceDollars      string      `json:"yes_price_dollars"`
	NoPriceDollars       string      `json:"no_price_dollars"`
	CreatedTime          time.Time   `json:"created_time"`
	LastUpdateTime       time.Time   `json:"last_update_time"`
	ExpirationTime       *time.Time  `json:"expiration_time,omitempty"`
	InitialCount         int         `json:"initial_count"`
	InitialCountFP       string      `json:"initial_count_fp"`
	RemainingCount       int         `json:"remaining_count"`
	RemainingCountFP     string      `json:"remaining_count_fp"`
	FillCount            int         `json:"fill_count"`
	FillCountFP          string      `json:"fill_count_fp"`
	TakerFillCount       int         `json:"taker_fill_count"`
	TakerFillCost        int         `json:"taker_fill_cost"`
	TakerFillCostDollars string      `json:"taker_fill_cost_dollars"`
	MakerFillCount       int         `json:"maker_fill_count"`
	MakerFillCost        int         `json:"maker_fill_cost"`
	MakerFillCostDollars string      `json:"maker_fill_cost_dollars"`
	TakerFees            int         `json:"taker_fees"`
	TakerFeesDollars     string      `json:"taker_fees_dollars"`
	MakerFees            int         `json:"maker_fees"`
	MakerFeesDollars     string      `json:"maker_fees_dollars"`
	QueuePosition        int         `json:"queue_position"`
	OrderGroupID         string      `json:"order_group_id"`
	CancelOrderOnPause   bool        `json:"cancel_order_on_pause"`
	SelfTradePreventType string      `json:"self_trade_prevention_type"`
}

// TotalFillCount returns the total number of filled contracts.
func (o *Order) TotalFillCount() int {
	return o.TakerFillCount + o.MakerFillCount
}

// TotalFillCost returns the total cost of filled contracts.
func (o *Order) TotalFillCost() int {
	return o.TakerFillCost + o.MakerFillCost
}

// IsFilled returns whether the order has been fully filled.
func (o *Order) IsFilled() bool {
	return o.Status == OrderStatusExecuted
}

// IsCanceled returns whether the order has been canceled.
func (o *Order) IsCanceled() bool {
	return o.Status == OrderStatusCanceled
}

// IsResting returns whether the order is resting in the order book.
func (o *Order) IsResting() bool {
	return o.Status == OrderStatusResting
}

// Position represents a user's position in a market.
type Position struct {
	Ticker             string `json:"ticker"`
	EventTicker        string `json:"event_ticker"`
	Position           int    `json:"position"` // positive = yes, negative = no
	TotalCost          int64  `json:"total_cost"`
	RealizedPnl        int64  `json:"realized_pnl"`
	RestingOrdersCount int    `json:"resting_orders_count"`
}

// Balance represents account balance information.
type Balance struct {
	Balance        int64 `json:"balance"`         // Available balance in cents
	PortfolioValue int64 `json:"portfolio_value"` // Portfolio value in cents
}

// Fill represents a trade fill.
type Fill struct {
	TradeID     string    `json:"trade_id"`
	Ticker      string    `json:"ticker"`
	Side        Side      `json:"side"`
	Action      Action    `json:"action"`
	Count       int       `json:"count"`
	YesPrice    int       `json:"yes_price"`
	NoPrice     int       `json:"no_price"`
	IsTaker     bool      `json:"is_taker"`
	OrderID     string    `json:"order_id"`
	CreatedTime time.Time `json:"created_time"`
}

// Settlement represents a market settlement.
type Settlement struct {
	Ticker       string    `json:"ticker"`
	MarketResult string    `json:"market_result"`
	YesCount     int       `json:"yes_count"`
	NoCount      int       `json:"no_count"`
	YesCost      int64     `json:"yes_cost"`
	NoCost       int64     `json:"no_cost"`
	Revenue      int64     `json:"revenue"`
	SettledTime  time.Time `json:"settled_time"`
}

// Orderbook represents the order book for a market.
// Note: Kalshi returns yes/no as arrays of [price, count] pairs
type Orderbook struct {
	Ticker     string  `json:"ticker"`
	Yes        [][]int `json:"yes"`         // [[price, count], ...] - YES bids (legacy integer)
	No         [][]int `json:"no"`          // [[price, count], ...] - NO bids (legacy integer)
	YesDollars [][]any `json:"yes_dollars"` // [["price_str", count], ...] - YES bids in dollars
	NoDollars  [][]any `json:"no_dollars"`  // [["price_str", count], ...] - NO bids in dollars
}

// OrderbookLevel represents a parsed price level in the order book.
type OrderbookLevel struct {
	Price int // Price in cents
	Count int // Number of contracts
}

// GetYesLevels parses the Yes orderbook into OrderbookLevel slice
func (o *Orderbook) GetYesLevels() []OrderbookLevel {
	levels := make([]OrderbookLevel, 0, len(o.Yes))
	for _, pair := range o.Yes {
		if len(pair) >= 2 {
			levels = append(levels, OrderbookLevel{Price: pair[0], Count: pair[1]})
		}
	}
	return levels
}

// GetNoLevels parses the No orderbook into OrderbookLevel slice
func (o *Orderbook) GetNoLevels() []OrderbookLevel {
	levels := make([]OrderbookLevel, 0, len(o.No))
	for _, pair := range o.No {
		if len(pair) >= 2 {
			levels = append(levels, OrderbookLevel{Price: pair[0], Count: pair[1]})
		}
	}
	return levels
}

// CreateOrderParams contains parameters for creating an order.
type CreateOrderParams struct {
	Ticker        string      `json:"ticker"`
	Side          Side        `json:"side"`
	Action        Action      `json:"action"`
	Count         int         `json:"count"`
	Type          OrderType   `json:"type,omitempty"`
	YesPrice      int         `json:"yes_price,omitempty"`
	NoPrice       int         `json:"no_price,omitempty"`
	ClientOrderID string      `json:"client_order_id,omitempty"`
	TimeInForce   TimeInForce `json:"time_in_force,omitempty"`
	ExpirationTS  int64       `json:"expiration_ts,omitempty"`
	BuyMaxCost    int         `json:"buy_max_cost,omitempty"` // Maximum cost in cents for market orders (auto FoK)
}

// ListParams contains common pagination parameters.
type ListParams struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// ListMarketsParams contains parameters for listing markets.
type ListMarketsParams struct {
	ListParams
	EventTicker string       `json:"event_ticker,omitempty"`
	Status      MarketStatus `json:"status,omitempty"`
	Tickers     []string     `json:"tickers,omitempty"`
}

// ListOrdersParams contains parameters for listing orders.
type ListOrdersParams struct {
	ListParams
	Ticker string      `json:"ticker,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
}
