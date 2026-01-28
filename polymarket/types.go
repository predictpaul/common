// Package polymarket provides common types for the Polymarket prediction market API.
// This package can be imported by third-party applications via:
//
//	go get github.com/predictpaul/common/polymarket
package polymarket

// OrderStatus order status enumeration
type OrderStatus string

const (
	OrderStatusPENDING   OrderStatus = "PENDING"
	OrderStatusMATCHED   OrderStatus = "MATCHED"
	OrderStatusUNMATCHED OrderStatus = "UNMATCHED"
	OrderStatusLIVE      OrderStatus = "LIVE"
	OrderStatusDELAY     OrderStatus = "DELAY"
	OrderStatusCANCELED  OrderStatus = "CANCELED"
)

// PolymarketEvent represents market event data
type PolymarketEvent struct {
	ID                    string  `json:"id"`
	Ticker                string  `json:"ticker"`
	Slug                  string  `json:"slug"`
	Title                 string  `json:"title"`
	Description           string  `json:"description"`
	ResolutionSource      string  `json:"resolutionSource"`
	StartDate             string  `json:"startDate"`
	CreationDate          string  `json:"creationDate"`
	EndDate               string  `json:"endDate"`
	Image                 string  `json:"image"`
	Icon                  string  `json:"icon"`
	Active                bool    `json:"active"`
	Closed                bool    `json:"closed"`
	Archived              bool    `json:"archived"`
	New                   bool    `json:"new"`
	Featured              bool    `json:"featured"`
	Restricted            bool    `json:"restricted"`
	Volume                float64 `json:"volume"`
	OpenInterest          float64 `json:"openInterest"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
	Volume1wk             float64 `json:"volume1wk"`
	Volume1mo             float64 `json:"volume1mo"`
	Volume1yr             float64 `json:"volume1yr"`
	EnableOrderBook       bool    `json:"enableOrderBook"`
	NegRisk               bool    `json:"negRisk"`
	CommentCount          int     `json:"commentCount"`
	Cyom                  bool    `json:"cyom"`
	ClosedTime            string  `json:"closedTime"`
	ShowAllOutcomes       bool    `json:"showAllOutcomes"`
	ShowMarketImages      bool    `json:"showMarketImages"`
	AutomaticallyResolved bool    `json:"automaticallyResolved"`
	EnableNegRisk         bool    `json:"enableNegRisk"`
	AutomaticallyActive   bool    `json:"automaticallyActive"`
	GmpChartMode          string  `json:"gmpChartMode"`
	NegRiskAugmented      bool    `json:"negRiskAugmented"`
	PendingDeployment     bool    `json:"pendingDeployment"`
	Deploying             bool    `json:"deploying"`
	RequiresTranslation   bool    `json:"requiresTranslation"`
}

// ClobReward represents market reward structure
type ClobReward struct {
	ID               string  `json:"id"`
	ConditionID      string  `json:"conditionId"`
	AssetAddress     string  `json:"assetAddress"`
	RewardsAmount    float64 `json:"rewardsAmount"`
	RewardsDailyRate float64 `json:"rewardsDailyRate"`
	StartDate        string  `json:"startDate"`
	EndDate          string  `json:"endDate"`
}

// PolymarketMarket represents Polymarket market data
type PolymarketMarket struct {
	ID                           string            `json:"id"`
	Question                     string            `json:"question"`
	ConditionID                  string            `json:"conditionId"`
	Slug                         string            `json:"slug"`
	ResolutionSource             string            `json:"resolutionSource"`
	EndDate                      string            `json:"endDate"`
	StartDate                    string            `json:"startDate"`
	Image                        string            `json:"image"`
	Icon                         string            `json:"icon"`
	Description                  string            `json:"description"`
	Outcomes                     string            `json:"outcomes"`      // JSON string array
	OutcomePrices                string            `json:"outcomePrices"` // JSON string array
	Volume                       string            `json:"volume"`        // String representation of float
	Active                       bool              `json:"active"`
	Closed                       bool              `json:"closed"`
	MarketMakerAddress           string            `json:"marketMakerAddress"`
	CreatedAt                    string            `json:"createdAt"`
	UpdatedAt                    string            `json:"updatedAt"`
	ClosedTime                   string            `json:"closedTime"`
	New                          bool              `json:"new"`
	Featured                     bool              `json:"featured"`
	SubmittedBy                  string            `json:"submitted_by"`
	Archived                     bool              `json:"archived"`
	ResolvedBy                   string            `json:"resolvedBy"`
	Restricted                   bool              `json:"restricted"`
	GroupItemTitle               string            `json:"groupItemTitle"`
	GroupItemThreshold           string            `json:"groupItemThreshold"`
	QuestionID                   string            `json:"questionID"`
	UmaEndDate                   string            `json:"umaEndDate"`
	EnableOrderBook              bool              `json:"enableOrderBook"`
	OrderPriceMinTickSize        float64           `json:"orderPriceMinTickSize"`
	OrderMinSize                 float64           `json:"orderMinSize"`
	UmaResolutionStatus          string            `json:"umaResolutionStatus"`
	VolumeNum                    float64           `json:"volumeNum"`
	EndDateIso                   string            `json:"endDateIso"`
	StartDateIso                 string            `json:"startDateIso"`
	HasReviewedDates             bool              `json:"hasReviewedDates"`
	Volume1wk                    float64           `json:"volume1wk"`
	Volume1mo                    float64           `json:"volume1mo"`
	Volume1yr                    float64           `json:"volume1yr"`
	ClobTokenIds                 string            `json:"clobTokenIds"` // JSON string array
	UmaBond                      string            `json:"umaBond"`
	UmaReward                    string            `json:"umaReward"`
	Volume1wkClob                float64           `json:"volume1wkClob"`
	Volume1moClob                float64           `json:"volume1moClob"`
	Volume1yrClob                float64           `json:"volume1yrClob"`
	VolumeClob                   float64           `json:"volumeClob"`
	CustomLiveness               float64           `json:"customLiveness"`
	AcceptingOrders              bool              `json:"acceptingOrders"`
	NegRisk                      bool              `json:"negRisk"`
	NegRiskRequestID             string            `json:"negRiskRequestID"`
	Events                       []PolymarketEvent `json:"events"`
	Ready                        bool              `json:"ready"`
	Funded                       bool              `json:"funded"`
	AcceptingOrdersTimestamp     string            `json:"acceptingOrdersTimestamp"`
	Cyom                         bool              `json:"cyom"`
	PagerDutyNotificationEnabled bool              `json:"pagerDutyNotificationEnabled"`
	Approved                     bool              `json:"approved"`
	ClobRewards                  []ClobReward      `json:"clobRewards"`
	RewardsMinSize               float64           `json:"rewardsMinSize"`
	RewardsMaxSpread             float64           `json:"rewardsMaxSpread"`
	Spread                       float64           `json:"spread"`
	AutomaticallyResolved        bool              `json:"automaticallyResolved"`
	LastTradePrice               float64           `json:"lastTradePrice"`
	BestBid                      float64           `json:"bestBid"`
	BestAsk                      float64           `json:"bestAsk"`
	AutomaticallyActive          bool              `json:"automaticallyActive"`
	ClearBookOnStart             bool              `json:"clearBookOnStart"`
	SeriesColor                  string            `json:"seriesColor"`
	ShowGmpSeries                bool              `json:"showGmpSeries"`
	ShowGmpOutcome               bool              `json:"showGmpOutcome"`
	ManualActivation             bool              `json:"manualActivation"`
	NegRiskOther                 bool              `json:"negRiskOther"`
	UmaResolutionStatuses        string            `json:"umaResolutionStatuses"` // JSON string array
	PendingDeployment            bool              `json:"pendingDeployment"`
	Deploying                    bool              `json:"deploying"`
	DeployingTimestamp           string            `json:"deployingTimestamp"`
	RfqEnabled                   bool              `json:"rfqEnabled"`
	HoldingRewardsEnabled        bool              `json:"holdingRewardsEnabled"`
	FeesEnabled                  bool              `json:"feesEnabled"`
	RequiresTranslation          bool              `json:"requiresTranslation"`
}

// PolymarketResponse represents returned market data array
type PolymarketResponse []PolymarketMarket

// ClobToken represents CLOB API token info
type ClobToken struct {
	TokenID string  `json:"token_id"`
	Outcome string  `json:"outcome"`
	Price   float64 `json:"price"`
	Winner  bool    `json:"winner"`
}

// ClobRewardsInfo represents CLOB API rewards info
type ClobRewardsInfo struct {
	Rates     interface{} `json:"rates"`
	MinSize   float64     `json:"min_size"`
	MaxSpread float64     `json:"max_spread"`
}

// ClobMarket represents CLOB API /markets/{condition_id} market data
type ClobMarket struct {
	EnableOrderBook         bool            `json:"enable_order_book"`
	Active                  bool            `json:"active"`
	Closed                  bool            `json:"closed"`
	Archived                bool            `json:"archived"`
	AcceptingOrders         bool            `json:"accepting_orders"`
	AcceptingOrderTimestamp string          `json:"accepting_order_timestamp"`
	MinimumOrderSize        float64         `json:"minimum_order_size"`
	MinimumTickSize         float64         `json:"minimum_tick_size"`
	ConditionID             string          `json:"condition_id"`
	QuestionID              string          `json:"question_id"`
	Question                string          `json:"question"`
	Description             string          `json:"description"`
	MarketSlug              string          `json:"market_slug"`
	EndDateIso              string          `json:"end_date_iso"`
	GameStartTime           *string         `json:"game_start_time"`
	SecondsDelay            int             `json:"seconds_delay"`
	Fpmm                    string          `json:"fpmm"`
	MakerBaseFee            float64         `json:"maker_base_fee"`
	TakerBaseFee            float64         `json:"taker_base_fee"`
	NotificationsEnabled    bool            `json:"notifications_enabled"`
	NegRisk                 bool            `json:"neg_risk"`
	NegRiskMarketID         string          `json:"neg_risk_market_id"`
	NegRiskRequestID        string          `json:"neg_risk_request_id"`
	Icon                    string          `json:"icon"`
	Image                   string          `json:"image"`
	Rewards                 ClobRewardsInfo `json:"rewards"`
	Is5050Outcome           bool            `json:"is_50_50_outcome"`
	Tokens                  []ClobToken     `json:"tokens"`
	Tags                    []string        `json:"tags"`
}

// IsClosed returns whether the market is closed (not accepting orders).
func (m *ClobMarket) IsClosed() bool {
	return m.Closed || !m.AcceptingOrders
}

// IsSettled returns whether the market is settled (has a winner).
func (m *ClobMarket) IsSettled() bool {
	if !m.Closed {
		return false
	}
	// Check if any token has a winner
	for _, token := range m.Tokens {
		if token.Winner {
			return true
		}
	}
	return false
}

// GetUnifiedStatus returns the unified market status.
// Polymarket status mapping:
//   - active=true, closed=false -> open
//   - closed=true, no winner -> closed
//   - closed=true, has winner -> settled
func (m *ClobMarket) GetUnifiedStatus() string {
	if !m.Closed && m.Active {
		return "open"
	}
	if m.IsSettled() {
		return "settled"
	}
	return "closed"
}

// GetResult returns the winning outcome ("Yes" or "No"), or empty if not settled.
func (m *ClobMarket) GetResult() string {
	for _, token := range m.Tokens {
		if token.Winner {
			return token.Outcome
		}
	}
	return ""
}

// TradeParams trade query parameters (matches official py-clob-client)
type TradeParams struct {
	ID      string `json:"id,omitempty"`
	Maker   string `json:"maker_address,omitempty"` // funder/platform address to get trades for
	Market  string `json:"market,omitempty"`        // condition ID
	AssetId string `json:"asset_id,omitempty"`      // token ID
	Before  int64  `json:"before,omitempty"`        // unix timestamp
	After   int64  `json:"after,omitempty"`         // unix timestamp
}

// MakerOrder maker order info
type MakerOrder struct {
	OrderID       string `json:"order_id"`
	Owner         string `json:"owner"`
	MakerAddress  string `json:"maker_address"`
	MatchedAmount string `json:"matched_amount"`
	Price         string `json:"price"`
	FeeRateBps    string `json:"fee_rate_bps"`
	AssetID       string `json:"asset_id"`
	Outcome       string `json:"outcome"`
	Side          string `json:"side"`
}

// Trade trade record
// Status: MATCHED(processing) -> MINED(on-chain) -> CONFIRMED(final state-success) / RETRYING(retrying) / FAILED(final state-failed)
type Trade struct {
	ID              string       `json:"id"`
	TakerOrderID    string       `json:"taker_order_id"`
	Market          string       `json:"market"`   // condition id
	AssetID         string       `json:"asset_id"` // token id
	Side            string       `json:"side"`     // BUY or SELL
	Size            string       `json:"size"`
	FeeRateBps      string       `json:"fee_rate_bps"`
	Price           string       `json:"price"`
	Status          string       `json:"status"` // MATCHED/MINED/CONFIRMED/RETRYING/FAILED
	MatchTime       string       `json:"match_time"`
	LastUpdate      string       `json:"last_update"`
	Outcome         string       `json:"outcome"`
	BucketIndex     int          `json:"bucket_index"`
	Owner           string       `json:"owner"`         // api key of taker
	MakerAddress    string       `json:"maker_address"` // funder address of taker
	TransactionHash string       `json:"transaction_hash"`
	MakerOrders     []MakerOrder `json:"maker_orders"`
	Type            string       `json:"type"` // TAKER or MAKER
}

// TradesResponse trade list response
type TradesResponse struct {
	NextCursor string  `json:"next_cursor"`
	Data       []Trade `json:"data"`
}
