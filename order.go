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

// OrderCreateResponse represents the response for POST /order/create.
type OrderCreateResponse struct {
	ID              string     `json:"id"`
	UserWallet      string     `json:"user_wallet"`
	MarketType      string     `json:"market_type"`
	MarketID        string     `json:"market_id"`
	TokenID         string     `json:"token_id"`
	EventID         string     `json:"event_id,omitempty"`
	MarketSide      string     `json:"market_side"`
	OrderDirection  string     `json:"order_direction"`
	OrderType       string     `json:"order_type"`
	TokenAmount     string     `json:"token_amount"`
	SharesAmount    string     `json:"shares_amount"`
	FilledCost      string     `json:"filled_cost"`
	FilledPrice     string     `json:"filled_price"`
	FeesPaid        string     `json:"fees_paid"`
	Status          string     `json:"status"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`
}

// OpenOrderQuery represents the request for GET|POST /order/open.
type OpenOrderQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet"`
	EventID    string `json:"event_id,omitempty" form:"event_id"`
	Page       int    `json:"page,omitempty" form:"page"`
	PageSize   int    `json:"page_size,omitempty" form:"page_size"`
}

// OpenOrderItem represents a single open order.
type OpenOrderItem struct {
	OrderID   string     `json:"order_id"`
	EventID   string     `json:"event_id"`
	TokenID   string     `json:"token_id"`
	MarketID  string     `json:"market_id"`
	Side      string     `json:"side"`      // YES / NO
	Type      string     `json:"type"`      // BUY / SELL
	Price     string     `json:"price"`
	Filled    string     `json:"filled"`
	Total     string     `json:"total"`
	Status    string     `json:"status"`
	Source    string     `json:"source"`    // POLYMARKET / KALSHI
	OrderTime *time.Time `json:"order_time"`
}

// OpenOrderResponse represents the response for GET|POST /order/open.
type OpenOrderResponse struct {
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Orders   []OpenOrderItem `json:"orders"`
}

// OrderHistoryQuery represents the request for GET|POST /order/history.
type OrderHistoryQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet"`
	EventID    string `json:"event_id,omitempty" form:"event_id"`
	Page       int    `json:"page,omitempty" form:"page"`
	PageSize   int    `json:"page_size,omitempty" form:"page_size"`
}

// OrderHistoryItem represents a single order in history.
type OrderHistoryItem struct {
	OrderID   string `json:"order_id"`
	CreatedAt string `json:"created_at"`
	Type      string `json:"type"`      // BUY / SELL
	Source    string `json:"source"`    // POLYMARKET / KALSHI
	EventID   string `json:"event_id"`
	Side      string `json:"side"`      // YES / NO
	Price     string `json:"price"`
	Shares    string `json:"shares"`
	Filled    string `json:"filled"`
	Cost      string `json:"cost"`
	Change    string `json:"change"`
	Status    string `json:"status"`
}

// OrderHistoryResponse represents the response for GET|POST /order/history.
type OrderHistoryResponse struct {
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Orders   []OrderHistoryItem `json:"orders"`
}

// AdminOrderListQuery represents the request for GET|POST /order/list (admin).
type AdminOrderListQuery struct {
	PageNum   int                    `json:"pageNum,omitempty" form:"pageNum"`
	PageSize  int                    `json:"pageSize,omitempty" form:"pageSize"`
	Search    map[string]interface{} `json:"search,omitempty" form:"search"`
	SortField string                 `json:"sortField,omitempty" form:"sortField"`
	SortOrder string                 `json:"sortOrder,omitempty" form:"sortOrder"` // ascend / descend
}

// AdminOrderItem represents an order item in admin list.
type AdminOrderItem struct {
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
	StopPrice       string     `json:"stop_price"`
	TakeProfitPrice string     `json:"take_profit_price"`
	FilledCost      string     `json:"filled_cost"`
	FilledPrice     string     `json:"filled_price"`
	FeesPaid        string     `json:"fees_paid"`
	Status          string     `json:"status"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
