package common

import "time"

// ---- Market / Order enums ----

// MarketType represents the market platform type.
const (
	MarketTypePolymarket = "POLYMARKET"
	MarketTypeKalshi     = "KALSHI"
	MarketTypeOpinion    = "OPINION"
)

// MarketSide represents the market outcome side.
const (
	MarketSideYES = "YES"
	MarketSideNO  = "NO"
)

// OrderDirection represents the direction of an order.
const (
	OrderDirectionBUY  = "BUY"
	OrderDirectionSELL = "SELL"
)

// OrderType represents the type of an order.
const (
	OrderTypeMarket = "MARKET"
	OrderTypeLimit  = "LIMIT"
	OrderTypeStop   = "STOP"
)

// OrderStatus represents the status of an order.
// Terminal states: FILLED, CANCELLED, FAILED.
const (
	OrderStatusSubmitting    = "SUBMITTING"     // 订单已创建，资金已锁定，正在提交到平台
	OrderStatusPending       = "PENDING"        // 平台已接收，挂单中
	OrderStatusPartialFilled = "PARTIAL_FILLED" // 部分成交，订单仍在平台上
	OrderStatusCancelling    = "CANCELLING"     // 用户请求取消，等待后台执行
	OrderStatusFilled        = "FILLED"         // 已成交（终态）
	OrderStatusCancelled     = "CANCELLED"      // 已取消（终态）
	OrderStatusFailed        = "FAILED"         // 提交失败（终态）
)

// IsTerminalStatus returns true if the status is a terminal state (no further transitions).
func IsTerminalStatus(status string) bool {
	return status == OrderStatusFilled || status == OrderStatusCancelled || status == OrderStatusFailed
}

// IsCancellableStatus returns true if the status allows user-initiated cancellation.
func IsCancellableStatus(status string) bool {
	return status == OrderStatusSubmitting || status == OrderStatusPending || status == OrderStatusPartialFilled
}

// ---- Request types ----

// OrderCreateRequest represents a single order in a batch creation request.
type OrderCreateRequest struct {
	UserWallet      string   `json:"user_wallet"`
	MarketType      string   `json:"market_type"`       // POLYMARKET / KALSHI / OPINION
	TokenID         string   `json:"token_id"`
	MarketID        string   `json:"market_id"`
	EventID         string   `json:"event_id,omitempty"`
	MarketSide      string   `json:"market_side"`       // YES / NO
	OrderDirection  string   `json:"order_direction"`   // BUY / SELL
	OrderType       string   `json:"order_type"`        // MARKET / LIMIT / STOP
	TokenAmount     string   `json:"token_amount,omitempty"`
	LimitPrice      string   `json:"limit_price,omitempty"`
	SharesAmount    string   `json:"shares_amount,omitempty"`
	StopPrice       string   `json:"stop_price,omitempty"`
	TakeProfitPrice string   `json:"take_profit_price,omitempty"`
	IdempotencyKey  string   `json:"idempotency_key,omitempty"`
	FeesEnabled     bool     `json:"fees_enabled,omitempty"`      // 是否启用手续费（Polymarket 部分市场启用）
	MarketTags      []string `json:"market_tags,omitempty"`       // 市场标签，用于判断手续费类型 e.g. ["crypto","Bitcoin"] / ["sport","NCAAB"]
}

// BatchOrderCreateRequest represents a batch order creation request.
// Common fields (user_wallet, market_side, order_direction) are top-level;
// per-order fields are in the list array and inherit common fields if not set.
type BatchOrderCreateRequest struct {
	UserWallet     string               `json:"user_wallet"`
	MarketSide     string               `json:"market_side"`     // YES / NO
	OrderDirection string               `json:"order_direction"` // BUY / SELL
	IdempotencyKey string               `json:"idempotency_key,omitempty"`
	List           []OrderCreateRequest `json:"list"`
}

// OrderCancelRequest represents a request to cancel specific orders.
type OrderCancelRequest struct {
	UserWallet string   `json:"user_wallet"`
	IDList     []string `json:"id_list"`
}

// OrderCancelAllRequest represents a request to cancel all active orders for a user.
type OrderCancelAllRequest struct {
	UserWallet string `json:"user_wallet"`
}

// ---- Response types ----

// CancelResult represents the result of an async cancel operation.
type CancelResult struct {
	AcceptedIDs []string `json:"accepted_ids"` // Orders accepted for cancellation (now CANCELLING)
	RejectedIDs []string `json:"rejected_ids"` // Orders that cannot be cancelled (wrong status/owner/not found)
}

// OrderCreateResponse is the response for POST /order/create.
// Same as OrderItem — both endpoints return model.Order.
type OrderCreateResponse = OrderItem

// ---- Query types ----

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

// ---- Item types ----

// OrderItem represents an order returned by GET /order/list.
// Mirrors model.Order JSON serialization.
type OrderItem struct {
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
	RequestedAmount string     `json:"requested_amount"`
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
	FeesEnabled     bool       `json:"fees_enabled"`
	MarketTags      []string   `json:"market_tags"`
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

// ---- Event order types ----

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
	FeesEnabled     bool       `json:"fees_enabled"`
	MarketTags      []string   `json:"market_tags"`
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
