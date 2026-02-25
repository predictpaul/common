package common

import "github.com/shopspring/decimal"

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
	TokenID      string          `json:"token_id"`
	MarketID     string          `json:"market_id"`
	MarketType   string          `json:"market_type"`
	MarketSide   string          `json:"market_side"`
	Balance      decimal.Decimal `json:"balance"`
	CurrentPrice decimal.Decimal `json:"current_price"`
	CurrentValue decimal.Decimal `json:"current_value"`
	IsSettle     bool            `json:"is_settle"`
}

// BalanceResponse represents the response for GET /account/balance.
type BalanceResponse struct {
	USDC           decimal.Decimal `json:"usdc"`
	TotalValue     decimal.Decimal `json:"total_value"`
	TotalTokens    int             `json:"total_tokens"`
	TokenPositions []TokenPosition `json:"token_positions"`
}

// MarketFilter represents a market filter for positions and orders queries.
type MarketFilter struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}

// PositionQuery represents query parameters for GET /account/positions.
type PositionQuery struct {
	UserWallet string         `json:"user_wallet" form:"user_wallet"`
	Page       int            `json:"page,omitempty" form:"page"`
	PageSize   int            `json:"page_size,omitempty" form:"page_size"`
	Markets    []MarketFilter `json:"markets,omitempty"`
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
	TotalShares     string         `json:"total_shares"`
	TotalPnL        string         `json:"total_pnl"`
	TotalPnLPercent string         `json:"total_pnl_percent"`
	PositionCount   int            `json:"position_count"`
	Total           int64          `json:"total"`
	Page            int            `json:"page"`
	PageSize        int            `json:"page_size"`
	Positions       []PositionItem `json:"positions"`
}

// PortfolioQuery represents query parameters for GET /account/portfolio.
type PortfolioQuery struct {
	UserWallet string         `json:"user_wallet" form:"user_wallet"`
	Markets    []MarketFilter `json:"markets,omitempty"`
}

// PortfolioResponse represents the response for GET /account/portfolio.
type PortfolioResponse struct {
	TotalPortfolioValue  string `json:"total_portfolio_value"`
	USDCBalance          string `json:"usdc_balance"`
	PositionsValue       string `json:"positions_value"`
	TotalCost            string `json:"total_cost"`
	UnrealizedPnL        string `json:"unrealized_pnl"`
	UnrealizedPnLPercent string `json:"unrealized_pnl_percent"`
	MaxPotential         string `json:"max_potential"`
	PositionCount        int    `json:"position_count"`
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
