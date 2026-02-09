package polymarket

import "time"

// Order represents an order in the API response.
type Order struct {
	ID               string    `json:"id"`
	UserWallet       string    `json:"user_wallet"`
	OrderTime        time.Time `json:"order_time"`
	MarketType       string    `json:"market_type"`
	MarketAccountID  string    `json:"market_account_id,omitempty"`
	MarketAccountKey string    `json:"market_account_key,omitempty"`
	MarketID         string    `json:"market_id"`
	MarketOutID      string    `json:"market_out_id,omitempty"`
	EventID          string    `json:"event_id,omitempty"`
	MarketSide       string    `json:"market_side"`
	MarketOrderID    string    `json:"market_order_id,omitempty"`
	TokenID          string    `json:"token_id"`
	TokenAmount      string    `json:"token_amount"`
	OrderDirection   string    `json:"order_direction"`
	OrderType        string    `json:"order_type"`
	LimitPrice       string    `json:"limit_price,omitempty"`
	SharesAmount     string    `json:"shares_amount,omitempty"`
	StopPrice        string    `json:"stop_price,omitempty"`
	TakeProfitPrice  string    `json:"take_profit_price,omitempty"`
	FilledCost       string    `json:"filled_cost"`
	FilledPrice      string    `json:"filled_price"`
	FeesPaid         string    `json:"fees_paid"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

// OrderListResponse represents the order list response data.
type OrderListResponse struct {
	Total    int     `json:"total"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	Orders   []Order `json:"orders"`
}

// CancelBatchResponse represents the batch cancel response data.
type CancelBatchResponse struct {
	SuccessIDs []string `json:"success_ids"`
	FailedIDs  []string `json:"failed_ids"`
}

// SyncResponse represents the sync response data.
type SyncResponse struct {
	SyncedCount int    `json:"synced_count"`
	Message     string `json:"message"`
}

// Account represents a user account.
type Account struct {
	ID         string    `json:"id"`
	UserWallet string    `json:"user_wallet"`
	Type       string    `json:"type"`
	Balance    string    `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
}

// TokenAccount represents a token position account.
type TokenAccount struct {
	ID         string `json:"id,omitempty"`
	TokenID    string `json:"token_id"`
	MarketID   string `json:"market_id,omitempty"`
	MarketType string `json:"market_type,omitempty"`
	MarketSide string `json:"market_side,omitempty"`
	Balance    string `json:"balance"`
	IsSettle   bool   `json:"is_settle"`
}

// WithdrawResponse represents the withdraw response data.
type WithdrawResponse struct {
	ID         string    `json:"id"`
	UserWallet string    `json:"user_wallet"`
	Amount     string    `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

// SettleResult represents a single settle result.
type SettleResult struct {
	Token   TokenAccount `json:"token"`
	Success bool         `json:"success"`
	Message string       `json:"message"`
}

// BalanceResponse represents the balance response data.
type BalanceResponse struct {
	USDC           string         `json:"usdc"`
	TotalTokens    int            `json:"total_tokens"`
	TokenPositions []TokenAccount `json:"token_positions"`
}

// Position represents a token position with PnL info.
type Position struct {
	TokenID      string `json:"token_id"`
	MarketID     string `json:"market_id"`
	MarketType   string `json:"market_type"`
	MarketSide   string `json:"market_side"`
	Balance      string `json:"balance"`
	AvgPrice     string `json:"avg_price"`
	CurrentPrice string `json:"current_price"`
	Pnl          string `json:"pnl"`
	PnlPercent   string `json:"pnl_percent"`
	IsSettle     bool   `json:"is_settle"`
}
