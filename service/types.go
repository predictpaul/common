// Package service provides common request/response types for prediction market service layer.
// This package can be imported by third-party applications via:
//
//	go get github.com/predictpaul/common/service
package service

import (
	"time"

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
	Markets      []MarketFilter    `json:"-" form:"-"`
}

// EventOrdersQuery represents event orders query parameters
type EventOrdersQuery struct {
	EventID      string            `json:"event_id" form:"event_id"`
	UserWallet   string            `json:"user_wallet" form:"user_wallet"`
	StatusFilter OrderStatusFilter `json:"status_filter" form:"status_filter"`
	Page         int               `json:"page" form:"page"`
	PageSize     int               `json:"page_size" form:"page_size"`
}

// OpenOrderQuery represents open order query parameters
type OpenOrderQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet"`
	EventID    string `json:"event_id" form:"event_id"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"page_size" form:"page_size"`
}

// OrderHistoryQuery represents order history query parameters
type OrderHistoryQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet"`
	EventID    string `json:"event_id" form:"event_id"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"page_size" form:"page_size"`
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

// MarketFilter represents a market filter for positions and orders queries.
// ID maps to market_out_id (orders) / market_id (accounts).
type MarketFilter struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}

// PositionQuery represents position query parameters
type PositionQuery struct {
	UserWallet string         `json:"user_wallet" form:"user_wallet" validate:"required"`
	Markets    []MarketFilter `json:"-" form:"-"`
}

// PortfolioQuery represents portfolio query parameters
type PortfolioQuery struct {
	UserWallet string         `json:"user_wallet" form:"user_wallet"`
	Markets    []MarketFilter `json:"-" form:"-"`
}

// EventPnLQuery represents event PnL query parameters
type EventPnLQuery struct {
	UserWallet string `json:"user_wallet" form:"user_wallet"`
	EventID    string `json:"event_id" form:"event_id"`
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
	TotalValue     decimal.Decimal `json:"total_value"`
	TotalTokens    int             `json:"total_tokens"`
	TokenPositions []TokenPosition `json:"token_positions"`
}

// TokenPosition represents token position information
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

// UnifiedMarketStatus represents the unified status of a market across platforms
type UnifiedMarketStatus string

const (
	MarketStatusOpen    UnifiedMarketStatus = "open"
	MarketStatusClosed  UnifiedMarketStatus = "closed"
	MarketStatusSettled UnifiedMarketStatus = "settled"
)

// PositionItem represents single position information
type PositionItem struct {
	TokenID              string              `json:"token_id"`
	MarketID             string              `json:"market_id"`
	EventID              string              `json:"event_id"`
	EventTitle           string              `json:"event_title"`
	Source               string              `json:"source"`
	MarketType           string              `json:"market_type"`
	MarketSide           string              `json:"market_side"`
	Shares               decimal.Decimal     `json:"shares"`
	Balance              decimal.Decimal     `json:"balance"`
	AvgCost              decimal.Decimal     `json:"avg_cost"`
	TotalCost            decimal.Decimal     `json:"total_cost"`
	CurrentPrice         decimal.Decimal     `json:"current_price"`
	CurrentValue         decimal.Decimal     `json:"current_value"`
	UnrealizedPnL        decimal.Decimal     `json:"unrealized_pnl"`
	UnrealizedPnLPercent decimal.Decimal     `json:"unrealized_pnl_percent"`
	IsSettle             bool                `json:"is_settle"`
	MarketStatus         UnifiedMarketStatus `json:"market_status"`
	MarketResult         string              `json:"market_result"`
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

// PortfolioResponse represents portfolio overview response
type PortfolioResponse struct {
	TotalPortfolioValue decimal.Decimal `json:"total_portfolio_value"`
	USDCBalance         decimal.Decimal `json:"usdc_balance"`
	PositionsValue      decimal.Decimal `json:"positions_value"`
	TotalCost           decimal.Decimal `json:"total_cost"`
	TotalPnL            decimal.Decimal `json:"total_pnl"`
	TotalPnLPercent     decimal.Decimal `json:"total_pnl_percent"`
	PositionCount       int             `json:"position_count"`
}

// EventPnLResponse represents event PnL response
type EventPnLResponse struct {
	EventID       string          `json:"event_id"`
	TotalCost     decimal.Decimal `json:"total_cost"`
	CurrentValue  decimal.Decimal `json:"current_value"`
	MaxProfit     decimal.Decimal `json:"max_profit"`
	UnrealizedPnL decimal.Decimal `json:"unrealized_pnl"`
	PnLPercent    decimal.Decimal `json:"pnl_percent"`
	Positions     []PositionItem  `json:"positions"`
}

// EventOrderItem represents an enriched order with price/cost/pnl info
type EventOrderItem struct {
	ID              string          `json:"id"`
	UserWallet      string          `json:"user_wallet"`
	OrderTime       *time.Time      `json:"order_time"`
	MarketType      string          `json:"market_type"`
	MarketID        string          `json:"market_id"`
	MarketOutID     string          `json:"market_out_id"`
	EventID         string          `json:"event_id"`
	MarketSide      string          `json:"market_side"`
	TokenID         string          `json:"token_id"`
	TokenAmount     decimal.Decimal `json:"token_amount"`
	OrderDirection  string          `json:"order_direction"`
	OrderType       string          `json:"order_type"`
	LimitPrice      decimal.Decimal `json:"limit_price"`
	RequestedShares decimal.Decimal `json:"requested_shares"`
	SharesAmount    decimal.Decimal `json:"shares_amount"`
	StopPrice       decimal.Decimal `json:"stop_price"`
	TakeProfitPrice decimal.Decimal `json:"take_profit_price"`
	FilledCost      decimal.Decimal `json:"filled_cost"`
	FilledPrice     decimal.Decimal `json:"filled_price"`
	FeesPaid        decimal.Decimal `json:"fees_paid"`
	Status          string          `json:"status"`
	CurrentPrice    decimal.Decimal `json:"current_price"`
	AvgCost         decimal.Decimal `json:"avg_cost"`
	CurrentValue    decimal.Decimal `json:"current_value"`
	PnL             decimal.Decimal `json:"pnl"`
	PnLPercent      decimal.Decimal `json:"pnl_percent"`
	Source          string          `json:"source"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// EventOrdersResponse represents event orders response with enriched data
type EventOrdersResponse struct {
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
	Orders   []EventOrderItem `json:"orders"`
}

// OpenOrderItem represents a single open order
type OpenOrderItem struct {
	OrderID   string          `json:"order_id"`
	EventID   string          `json:"event_id"`
	TokenID   string          `json:"token_id"`
	MarketID  string          `json:"market_id"`
	Side      string          `json:"side"`
	Type      string          `json:"type"`
	Price     decimal.Decimal `json:"price"`
	Filled    decimal.Decimal `json:"filled"`
	Total     decimal.Decimal `json:"total"`
	Status    string          `json:"status"`
	Source    string          `json:"source"`
	OrderTime *time.Time      `json:"order_time"`
}

// OpenOrderResponse represents open order response
type OpenOrderResponse struct {
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Orders   []OpenOrderItem `json:"orders"`
}

// OrderHistoryItem represents a single order in history
type OrderHistoryItem struct {
	OrderID   string          `json:"order_id"`
	CreatedAt string          `json:"created_at"`
	Type      string          `json:"type"`
	OrderType string          `json:"order_type"`
	Source    string          `json:"source"`
	EventID   string          `json:"event_id"`
	Side      string          `json:"side"`
	Price     decimal.Decimal `json:"price"`
	Shares    decimal.Decimal `json:"shares"`
	Filled    decimal.Decimal `json:"filled"`
	Cost      decimal.Decimal `json:"cost"`
	Change    decimal.Decimal `json:"change"`
	Status    string          `json:"status"`
}

// OrderHistoryResponse represents order history response
type OrderHistoryResponse struct {
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Orders   []OrderHistoryItem `json:"orders"`
}

// OrderListResponse represents order list response
type OrderListResponse struct {
	Total    int64  `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Orders   []any  `json:"orders"`
}

// AccountListResponse represents account list response
type AccountListResponse struct {
	Total    int64  `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Accounts []any  `json:"accounts"`
}

// FlowListResponse represents flow list response
type FlowListResponse struct {
	Total    int64  `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Flows    []any  `json:"flows"`
}

