// Package admin provides common types for PredictPaul admin API.
package admin

import "time"

// Node represents a blockchain node configuration.
type Node struct {
	ID            int       `json:"id"`
	ChainType     string    `json:"chain_type"`
	URL           string    `json:"url"`
	ChainID       int       `json:"chain_id"`
	ChainDecimals int       `json:"chain_decimals"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

// Wallet represents a wallet configuration.
type Wallet struct {
	ID            int       `json:"id"`
	ChainType     string    `json:"chain_type"`
	ChainID       int       `json:"chain_id"`
	WalletAddress string    `json:"wallet_address"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

// MarketConfig represents a market account configuration.
type MarketConfig struct {
	ID              string    `json:"id"`
	MarketType      string    `json:"market_type"`
	MarketSide      string    `json:"market_side"`
	APIKey          string    `json:"api_key"`
	PlatformAddress string    `json:"platform_address"`
	WalletAddress   string    `json:"wallet_address"`
	IsActive        int       `json:"is_active"`
	BaseURL         string    `json:"base_url"`
	ChainID         int       `json:"chain_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

// Node status
const (
	NodeStatusDisabled = 0
	NodeStatusEnabled  = 1
)

// Wallet status
const (
	WalletStatusDisabled = 0
	WalletStatusEnabled  = 1
)

// Market config status
const (
	MarketConfigInactive = 0
	MarketConfigActive   = 1
)
