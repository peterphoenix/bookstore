package repository

import (
	"context"

	"bookstore/domain/order/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (r repository) CreateOrder(ctx context.Context, input model.CreateOrderInput) error {
	err := r.psql.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		_, err := sq.Insert(orderTableName).
			Columns(
				orderTable.ID,
				orderTable.UserID,
				orderTable.TotalAmount,
				orderTable.Status,
				orderTable.ShippingAddress,
				orderTable.City,
				orderTable.State,
				orderTable.ZipCode,
				orderTable.Country,
			).
			Values(
				input.ID,
				input.UserID,
				input.Total,
				model.OrderStatusCompleted, //ignore order flow for now
				input.ShippingAddress,
				input.City,
				input.State,
				input.ZipCode,
				input.Country,
			).
			PlaceholderFormat(sq.Dollar).
			RunWith(tx).
			ExecContext(ctx)

		if err != nil {
			return err
		}

		insertOrderItemQuery := sq.Insert(orderItemTableName).
			Columns(
				orderItemTable.OrderID,
				orderItemTable.BookID,
				orderItemTable.Quantity,
				orderItemTable.Price,
			).
			PlaceholderFormat(sq.Dollar)
		for _, item := range input.Items {
			insertOrderItemQuery = insertOrderItemQuery.
				Values(
					input.ID,
					item.BookID,
					item.Qty,
					item.Price,
				)
		}

		_, err = insertOrderItemQuery.
			RunWith(tx).
			ExecContext(ctx)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
