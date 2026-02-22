package common

import "time"

// OrderCreateRequest represents the request for POST /order/create.
type OrderCreateRequest struct {
	UserWallet      string `json:"user_wallet"`
	MarketType      string `json:"market_type"`       // POLYMARKET / KALSHI
	TokenID         string `json:"token_id"`
	MarketID        string `json:"market_id"`
	EventID         string `json:"event_id,omitempty"`
	MarketSide      string `json:"market_side"`       // YES / NO
	OrderDirection  string `json:"order_direction"`   // BUY / SELL
	OrderType       string `json:"order_type"`        // MARKET / LIMIT / STOP
	TokenAmount     string `json:"token_amount,omitempty"`
	LimitPrice      string `json:"limit_price,omitempty"`
	SharesAmount    string `json:"shares_amount,omitempty"`
	StopPrice       string `json:"stop_price,omitempty"`
	TakeProfitPrice string `json:"take_profit_price,omitempty"`
}

// OrderCreateResponse is the response for POST /order/create.
// Same as OrderItem — both endpoints return model.Order.
type OrderCreateResponse = OrderItem

// OrderCancelAllRequest represents the request for POST /order/cancel-all.
type OrderCancelAllRequest struct {
	UserWallet string `json:"user_wallet"`
}

// CancelResult represents batch cancellation result.
type CancelResult struct {
	SuccessIDs []string `json:"success_ids"`
	FailedIDs  []string `json:"failed_ids"`
}

// OrderStatusFilter represents order status filter type.
type OrderStatusFilter string

const (
	OrderFilterAll      OrderStatusFilter = "all"
	OrderFilterFilled   OrderStatusFilter = "filled"
	OrderFilterUnfilled OrderStatusFilter = "unfilled"
	OrderFilterCanceled OrderStatusFilter = "canceled"
	OrderFilterSettled  OrderStatusFilter = "settled"
)

// OrderListQuery represents query parameters for GET /order/list.
type OrderListQuery struct {
	UserWallet   string            `json:"user_wallet" form:"user_wallet"`
	StatusFilter OrderStatusFilter `json:"status_filter,omitempty" form:"status_filter"`
	MarketType   string            `json:"market_type,omitempty" form:"market_type"`
	MarketID     string            `json:"market_id,omitempty" form:"market_id"`
	EventID      string            `json:"event_id,omitempty" form:"event_id"`
	TokenID      string            `json:"token_id,omitempty" form:"token_id"`
	Page         int               `json:"page,omitempty" form:"page"`
	PageSize     int               `json:"page_size,omitempty" form:"page_size"`
}

// OrderItem represents an order returned by GET /order/list.
// Mirrors model.Order JSON serialization.
type OrderItem struct {
	ID              string     `json:"id"`
	UserWallet      string     `json:"user_wallet"`
	OrderTime       *time.Time `json:"order_time"`
	MarketType      string     `json:"market_type"`
	MarketAccountID string     `json:"market_account_id"`
	MarketID        string     `json:"market_id"`
	MarketOutID     string     `json:"market_out_id"`
	EventID         string     `json:"event_id"`
	MarketSide      string     `json:"market_side"`
	MarketOrderID   string     `json:"market_order_id"`
	TokenID         string     `json:"token_id"`
	TokenAmount     string     `json:"token_amount"`
	OrderDirection  string     `json:"order_direction"`
	OrderType       string     `json:"order_type"`
	LimitPrice      string     `json:"limit_price"`
	RequestedShares string     `json:"requested_shares"`
	SharesAmount    string     `json:"shares_amount"`
	StopPrice       string     `json:"stop_price"`
	TakeProfitPrice string     `json:"take_profit_price"`
	FilledCost      string     `json:"filled_cost"`
	FilledPrice     string     `json:"filled_price"`
	FeesPaid        string     `json:"fees_paid"`
	Status          string     `json:"status"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// OrderListResponse represents the response for GET /order/list.
type OrderListResponse struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Orders   []OrderItem `json:"orders"`
}

// EventOrderItem represents an enriched order with price/cost/pnl info.
type EventOrderItem struct {
	ID              string     `json:"id"`
	UserWallet      string     `json:"user_wallet"`
	OrderTime       *time.Time `json:"order_time"`
	MarketType      string     `json:"market_type"`
	MarketID        string     `json:"market_id"`
	MarketOutID     string     `json:"market_out_id"`
	EventID         string     `json:"event_id"`
	MarketSide      string     `json:"market_side"`
	TokenID         string     `json:"token_id"`
	TokenAmount     string     `json:"token_amount"`
	OrderDirection  string     `json:"order_direction"`
	OrderType       string     `json:"order_type"`
	LimitPrice      string     `json:"limit_price"`
	RequestedShares string     `json:"requested_shares"`
	SharesAmount    string     `json:"shares_amount"`
	FilledCost      string     `json:"filled_cost"`
	FilledPrice     string     `json:"filled_price"`
	FeesPaid        string     `json:"fees_paid"`
	Status          string     `json:"status"`
	CurrentPrice    string     `json:"current_price"`
	AvgCost         string     `json:"avg_cost"`
	CurrentValue    string     `json:"current_value"`
	PnL             string     `json:"pnl"`
	PnLPercent      string     `json:"pnl_percent"`
	Source          string     `json:"source"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// EventOrdersResponse represents the response for GET|POST /order/event-orders.
type EventOrdersResponse struct {
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
	Orders   []EventOrderItem `json:"orders"`
}
