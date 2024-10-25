package model

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	OrderStatusCompleted = "COMPLETED"
)

type Order struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Status    string          `json:"status"`
	Total     decimal.Decimal `json:"total"`
	CreatedAt time.Time       `json:"created_at"`
}

type OrderItemDetails struct {
	BookID string          `json:"book_id"`
	Qty    uint            `json:"qty"`
	Price  decimal.Decimal `json:"price"`
}

type OrderDetails struct {
	Order
	Items []OrderItemDetails `json:"items"`
}
