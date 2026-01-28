// Package common provides common types for the PredictPaul API.
package common

// Response represents the unified API response format.
type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// PageData represents paginated response data.
type PageData[T any] struct {
	Data     []T `json:"data"`
	Count    int `json:"count"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// Response codes
const (
	CodeSuccess      = 0
	CodeFailed       = 101
	CodeUnauthorized = 102
)
