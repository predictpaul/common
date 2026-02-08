package common

// RechargeRequest represents the request for POST /account/recharge.
type RechargeRequest struct {
	UserWallet    string `json:"user_wallet"`
	UserTxHash    string `json:"user_tx_hash"`
	ChainID       int    `json:"chain_id"`
	TokenSymbol   string `json:"token_symbol,omitempty"`
	TokenAmount   string `json:"token_amount"`
	TokenDecimals int    `json:"token_decimals"`
}

// WithdrawRequest represents the request for POST /account/withdraw.
type WithdrawRequest struct {
	UserWallet string `json:"user_wallet"`
	ChainID    int    `json:"chain_id"`
	Amount     string `json:"amount"`
}

// SettleRequest represents the request for POST /account/settle.
type SettleRequest struct {
	UserWallet string `json:"user_wallet"`
	TokenID    string `json:"token_id,omitempty"`
}

// SettleResult represents settlement result for a single token.
type SettleResult struct {
	TokenID string `json:"token_id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// TokenPosition represents a token position in balance response.
type TokenPosition struct {
	TokenID      string `json:"token_id"`
	MarketID     string `json:"market_id"`
	MarketType   string `json:"market_type"`
	MarketSide   string `json:"market_side"`
	Balance      string `json:"balance"`
	CurrentPrice string `json:"current_price"`
	CurrentValue string `json:"current_value"`
	IsSettle     bool   `json:"is_settle"`
}

// BalanceResponse represents the response for GET /account/balance.
type BalanceResponse struct {
	USDC           string          `json:"usdc"`
	TotalValue     string          `json:"total_value"`
	TotalTokens    int             `json:"total_tokens"`
	TokenPositions []TokenPosition `json:"token_positions"`
}

// PositionItem represents a single position with PnL info.
type PositionItem struct {
	TokenID              string `json:"token_id"`
	MarketID             string `json:"market_id"`
	EventID              string `json:"event_id"`
	EventTitle           string `json:"event_title"`
	Source               string `json:"source"`       // POLYMARKET / KALSHI
	MarketType           string `json:"market_type"`
	MarketSide           string `json:"market_side"`
	Shares               string `json:"shares"`
	Balance              string `json:"balance"`
	AvgCost              string `json:"avg_cost"`
	TotalCost            string `json:"total_cost"`
	CurrentPrice         string `json:"current_price"`
	CurrentValue         string `json:"current_value"`
	UnrealizedPnL        string `json:"unrealized_pnl"`
	UnrealizedPnLPercent string `json:"unrealized_pnl_percent"`
	IsSettle             bool   `json:"is_settle"`
	MarketStatus         string `json:"market_status"` // open, closed, settled
	MarketResult         string `json:"market_result"` // yes, no, or empty
}

// PositionResponse represents the response for GET /account/positions.
type PositionResponse struct {
	TotalValue      string         `json:"total_value"`
	TotalCost       string         `json:"total_cost"`
	TotalPnL        string         `json:"total_pnl"`
	TotalPnLPercent string         `json:"total_pnl_percent"`
	PositionCount   int            `json:"position_count"`
	Positions       []PositionItem `json:"positions"`
}

// PortfolioResponse represents the response for GET /account/portfolio.
type PortfolioResponse struct {
	TotalPortfolioValue string `json:"total_portfolio_value"`
	USDCBalance         string `json:"usdc_balance"`
	PositionsValue      string `json:"positions_value"`
	TotalCost           string `json:"total_cost"`
	TotalPnL            string `json:"total_pnl"`
	TotalPnLPercent     string `json:"total_pnl_percent"`
	PositionCount       int    `json:"position_count"`
}

// EventPnLResponse represents the response for GET /account/event-pnl.
type EventPnLResponse struct {
	EventID       string         `json:"event_id"`
	TotalCost     string         `json:"total_cost"`
	CurrentValue  string         `json:"current_value"`
	MaxProfit     string         `json:"max_profit"`
	UnrealizedPnL string         `json:"unrealized_pnl"`
	PnLPercent    string         `json:"pnl_percent"`
	Positions     []PositionItem `json:"positions"`
}
