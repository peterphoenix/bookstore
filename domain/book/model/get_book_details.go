package model

import "github.com/shopspring/decimal"

type GetBookDetailsReq struct {
	IDs []string `json:"ids"`
}

type BookDetailsMap map[string]BookDetails

type BookDetails struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Price       decimal.Decimal `json:"price"`
	ImageURL    string          `json:"image_url"`
	ISBN        string          `json:"isbn"`
	Description string          `json:"description"`
}
