package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type GetOrderHistoryRes struct {
	OrderHistoryDetails []OrderHistoryDetail `json:"order_history_details"`
}

type OrderHistoryDetail struct {
	ID        string                    `json:"id"`
	UserID    string                    `json:"user_id"`
	Status    string                    `json:"status"`
	Total     decimal.Decimal           `json:"total"`
	CreatedAt time.Time                 `json:"created_at"`
	Items     []OrderHistoryItemDetails `json:"items"`
}

type OrderHistoryItemDetails struct {
	BookID   string          `json:"book_id"`
	Qty      uint            `json:"qty"`
	Price    decimal.Decimal `json:"price"`
	Title    string          `json:"title"`
	ImageURL string          `json:"image_url"`
}

type GetOrderHistoryFilter struct {
	StartTime time.Time
	EndTime   time.Time
}
