package model

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type CreateOrderItemReq struct {
	BookID string `json:"book_id"`
	Qty    uint   `json:"qty"`
}

type CreateOrderReq struct {
	Items           []CreateOrderItemReq `json:"items"`
	ShippingAddress string
	City            string
	State           string
	ZipCode         string
	Country         string
}

type CreateOrderRes struct {
	ID string `json:"id"`
}

type CreateOrderInput struct {
	ID              string
	Items           []OrderItemDetails
	UserID          string
	Total           decimal.Decimal
	ShippingAddress string
	City            string
	State           string
	ZipCode         string
	Country         string
}

func (req *CreateOrderReq) Validate() error {
	for _, item := range req.Items {
		if item.Qty == 0 {
			return errors.New("qty must be greater than zero")
		}
	}
	return nil
}