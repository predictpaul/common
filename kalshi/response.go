package kalshi

import "time"

// OrderResponse represents an order in the API response.
type OrderResponse struct {
	ID          int       `json:"id"`
	OrderID     string    `json:"order_id"`
	UserID      string    `json:"user_id"`
	Ticker      string    `json:"ticker"`
	Side        string    `json:"side"`
	Action      string    `json:"action"`
	OrderType   string    `json:"order_type,omitempty"`
	YesPrice    int       `json:"yes_price"`
	NoPrice     int       `json:"no_price,omitempty"`
	Count       int       `json:"count"`
	FilledCount int       `json:"filled_count"`
	FilledCost  int       `json:"filled_cost"`
	FeesPaid    int       `json:"fees_paid"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// AccountResponse represents a Kalshi account in the API response.
type AccountResponse struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   int64     `json:"balance"`
	Frozen    int64     `json:"frozen"`
	TotalPnl  int64     `json:"total_pnl"`
	CreatedAt time.Time `json:"created_at"`
}

// FlowResponse represents a fund flow record.
type FlowResponse struct {
	ID            int       `json:"id"`
	UserID        string    `json:"user_id"`
	OrderID       string    `json:"order_id,omitempty"`
	FlowType      int       `json:"flow_type"`
	Amount        int64     `json:"amount"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	Remark        string    `json:"remark,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

// MarketResponse represents a Kalshi market in the API response.
type MarketResponse struct {
	ID             int       `json:"id"`
	Ticker         string    `json:"ticker"`
	EventTicker    string    `json:"event_ticker"`
	Title          string    `json:"title"`
	Status         int       `json:"status"`
	Result         string    `json:"result"`
	YesBid         int       `json:"yes_bid"`
	YesAsk         int       `json:"yes_ask"`
	LastPrice      int       `json:"last_price"`
	Volume24H      int64     `json:"volume_24h"`
	OpenInterest   int64     `json:"open_interest"`
	CloseTime      time.Time `json:"close_time"`
	ExpirationTime time.Time `json:"expiration_time"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

// MarketListResponse represents the market list response data.
type MarketListResponse struct {
	Markets []MarketResponse `json:"markets"`
	Cursor  string           `json:"cursor,omitempty"`
}

// OrderbookResponse represents the orderbook response data.
type OrderbookResponse struct {
	Ticker string  `json:"ticker"`
	Yes    [][]int `json:"yes"` // [[price, count], ...]
	No     [][]int `json:"no"`  // [[price, count], ...]
}

// SettleResponse represents the market settle response data.
type SettleResponse struct {
	SettledCount int   `json:"settled_count"`
	TotalRevenue int64 `json:"total_revenue"`
}

// Flow types
const (
	FlowTypeRecharge = 1 // 充值
	FlowTypeWithdraw = 2 // 提现
	FlowTypeBuy      = 3 // 买入
	FlowTypeSell     = 4 // 卖出
	FlowTypeSettle   = 5 // 结算
	FlowTypeFee      = 6 // 手续费
	FlowTypeFreeze   = 7 // 冻结
	FlowTypeUnfreeze = 8 // 解冻
)

// API Order status (int values for internal API)
const (
	APIOrderStatusPending         = 0 // 待处理
	APIOrderStatusResting         = 1 // 挂单中
	APIOrderStatusFilled          = 2 // 已成交
	APIOrderStatusCanceled        = 3 // 已取消
	APIOrderStatusPartiallyFilled = 4 // 部分成交
	APIOrderStatusSettled         = 5 // 已结算
)
