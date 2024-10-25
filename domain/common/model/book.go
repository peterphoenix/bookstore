package model

import "github.com/shopspring/decimal"

// Book type
// If needed, can be added more fields such as quantity,
// published year and genre for filtering,
// dimensions and weight for shipment, etc.
type Book struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Authors     []string        `json:"authors"`
	Price       decimal.Decimal `json:"price"`
	ImageURL    string          `json:"image_url"`
	Description string          `json:"description"`
}
