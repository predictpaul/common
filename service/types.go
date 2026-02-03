// Package service provides common request/response types for prediction market service layer.
// This package can be imported by third-party applications via:
//
//	go get github.com/predictpaul/common/service
package service

import (
	"github.com/shopspring/decimal"
)

// =============================================================================
// Order Request Types
// =============================================================================

// OrderCreateRequest represents an order creation request
type OrderCreateRequest struct {
	UserWallet      string          `json:"user_wallet" validate:"required"`
	MarketType      string          `json:"market_type" validate:"required"`
	TokenID         string          `json:"token_id" validate:"required"`
	MarketID        string          `json:"market_id" validate:"required"`
	EventID         string          `json:"event_id"`
	MarketSide      string          `json:"market_side" validate:"required"`
	TokenAmount     decimal.Decimal `json:"token_amount"`
	OrderDirection  string          `json:"order_direction" validate:"required"`
	OrderType       string          `json:"order_type" validate:"required"`
	LimitPrice      decimal.Decimal `json:"limit_price"`
	SharesAmount    decimal.Decimal `json:"shares_amount"`
	StopPrice       decimal.Decimal `json:"stop_price"`
	TakeProfitPrice decimal.Decimal `json:"take_profit_price"`
}

// OrderCancelRequest represents an order cancellation request
type OrderCancelRequest struct {
	OrderID string `json:"order_id" validate:"required"`
}

// OrderCancelBatchRequest represents a batch order cancellation request
type OrderCancelBatchRequest struct {
	OrderIDs []string `json:"order_ids" validate:"required,min=1"`
}

// OrderCancelAllRequest represents a cancel all orders request
type OrderCancelAllRequest struct {
	UserWallet string `json:"user_wallet"`
}

// OrderCancelMarketRequest represents a cancel market orders request
type OrderCancelMarketRequest struct {
	UserWallet  string `json:"user_wallet"`
	MarketOutID string `json:"market_out_id"`
	TokenID     string `json:"token_id"`
}

// CancelResult represents batch cancellation result
type CancelResult struct {
	SuccessIDs []string `json:"success_ids"`
	FailedIDs  []string `json:"failed_ids"`
}

// =============================================================================
// Order Query Types
// =============================================================================

// OrderStatusFilter represents order status filter type
type OrderStatusFilter string

const (
	OrderFilterAll      OrderStatusFilter = "all"
	OrderFilterFilled   OrderStatusFilter = "filled"
	OrderFilterUnfilled OrderStatusFilter = "unfilled"
	OrderFilterCanceled OrderStatusFilter = "canceled"
	OrderFilterSettled  OrderStatusFilter = "settled"
)

// OrderListQuery represents order list query parameters
type OrderListQuery struct {
	UserWallet   string            `json:"user_wallet" form:"user_wallet" validate:"required"`
	StatusFilter OrderStatusFilter `json:"status_filter" form:"status_filter"`
	MarketType   string            `json:"market_type" form:"market_type"`
	MarketID     string            `json:"market_id" form:"market_id"`
	EventID      string            `json:"event_id" form:"event_id"`
	TokenID      string            `json:"token_id" form:"token_id"`
	Page         int               `json:"page" form:"page"`
	PageSize     int               `json:"page_size" form:"page_size"`
}

// =============================================================================
// Account Request Types
// =============================================================================

// RechargeRequest represents a recharge request
type RechargeRequest struct {
	UserWallet    string          `json:"user_wallet" validate:"required"`
	UserTxHash    string          `json:"user_tx_hash" validate:"required"`
	ChainID       int             `json:"chain_id" validate:"required"`
	TokenSymbol   string          `json:"token_symbol"`
	TokenAmount   decimal.Decimal `json:"token_amount" binding:"required"`
	TokenDecimals int             `json:"token_decimals" validate:"required"`
}

// WithdrawRequest represents a withdraw request
type WithdrawRequest struct {
	UserWallet string          `json:"user_wallet" validate:"required"`
	ChainID    int             `json:"chain_id" validate:"required"`
	Amount     decimal.Decimal `json:"amount" validate:"required"`
}

// SettleRequest represents a settle request
type SettleRequest struct {
	UserWallet string `json:"user_wallet" validate:"required"`
	TokenID    string `json:"token_id"`
}

// =============================================================================
// Account Query Types
// =============================================================================

// BalanceQuery represents balance query parameters
type BalanceQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet" validate:"required"`
}

// AccountListQuery represents account list query parameters
type AccountListQuery struct {
	UserWallet  string `json:"user_wallet" form:"user_wallet" validate:"required"`
	AccountType string `json:"type" form:"type"`
	IsSettle    *bool  `json:"is_settle" form:"is_settle"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"page_size" form:"page_size"`
}

// FlowListQuery represents flow list query parameters
type FlowListQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet" validate:"required"`
	Type       string `json:"type" form:"type"`
	Direction  string `json:"direction" form:"direction"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"page_size" form:"page_size"`
}

// PositionFilter represents a filter for specific position
type PositionFilter struct {
	MarketType string `json:"market_type"` // POLYMARKET or KALSHI
	EventID    string `json:"event_id"`
	TokenID    string `json:"token_id"`
}

// PositionQuery represents position query parameters
type PositionQuery struct {
	UserWallet string           `json:"user_wallet" form:"user_wallet" validate:"required"`
	Filters    []PositionFilter `json:"filters"` // Optional: if provided, only return matching positions
}

// =============================================================================
// Response Types
// =============================================================================

// SettleResult represents settlement result for a single token
type SettleResult struct {
	TokenID string `json:"token_id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// BalanceSummary represents user balance summary
type BalanceSummary struct {
	USDC           decimal.Decimal `json:"usdc"`
	TotalTokens    int             `json:"total_tokens"`
	TokenPositions []TokenPosition `json:"token_positions"`
}

// TokenPosition represents token position information
type TokenPosition struct {
	TokenID    string          `json:"token_id"`
	MarketID   string          `json:"market_id"`
	MarketType string          `json:"market_type"`
	MarketSide string          `json:"market_side"`
	Balance    decimal.Decimal `json:"balance"`
	IsSettle   bool            `json:"is_settle"`
}

// UnifiedMarketStatus represents the unified status of a market across platforms
type UnifiedMarketStatus string

const (
	MarketStatusOpen    UnifiedMarketStatus = "open"
	MarketStatusClosed  UnifiedMarketStatus = "closed"
	MarketStatusSettled UnifiedMarketStatus = "settled"
)

// PositionItem represents single position information
type PositionItem struct {
	TokenID      string              `json:"token_id"`
	MarketID     string              `json:"market_id"`
	MarketType   string              `json:"market_type"`
	MarketSide   string              `json:"market_side"`
	Balance      decimal.Decimal     `json:"balance"`
	AvgCost      decimal.Decimal     `json:"avg_cost"`
	TotalCost    decimal.Decimal     `json:"total_cost"`
	CurrentPrice decimal.Decimal     `json:"current_price"`
	CurrentValue decimal.Decimal     `json:"current_value"`
	PnL          decimal.Decimal     `json:"pnl"`
	PnLPercent   decimal.Decimal     `json:"pnl_percent"`
	IsSettle     bool                `json:"is_settle"`
	MarketStatus UnifiedMarketStatus `json:"market_status"`
	MarketResult string              `json:"market_result"`
}

// PositionResponse represents position response
type PositionResponse struct {
	TotalValue      decimal.Decimal `json:"total_value"`
	TotalCost       decimal.Decimal `json:"total_cost"`
	TotalPnL        decimal.Decimal `json:"total_pnl"`
	TotalPnLPercent decimal.Decimal `json:"total_pnl_percent"`
	PositionCount   int             `json:"position_count"`
	Positions       []PositionItem  `json:"positions"`
}

// CostInfo represents cost information
type CostInfo struct {
	TotalCost   decimal.Decimal
	TotalShares decimal.Decimal
	AvgCost     decimal.Decimal
}

// TransactionDetail represents transaction details
type TransactionDetail struct {
	Hash        string
	From        string
	To          string
	BlockNumber uint64
	GasUsed     uint64
	Status      uint64
}
