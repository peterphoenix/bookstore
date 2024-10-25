package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"bookstore/common/contextutil"
	"bookstore/common/pagination"
	"bookstore/domain/order/model"
	sq "github.com/Masterminds/squirrel"
)

func (r repository) GetOrderHistory(ctx context.Context, filter model.GetOrderHistoryFilter, pagination pagination.Pagination) ([]model.OrderDetails, error) {
	var orderDetails []model.OrderDetails
	userID := contextutil.GetUserID(ctx)

	jsonAggQuery := fmt.Sprintf(`
		JSON_AGG (
			JSON_BUILD_OBJECT (
				'%s', %s,
				'%s', %s,
				'%s', %s
			)
		) AS items
		`,
		orderItemTable.BookID, orderItemTable.BookID,
		orderItemTable.Quantity, orderItemTable.Quantity,
		orderItemTable.Price, orderItemTable.Price)

	query := sq.
		Select(
			orderTable.ID,
			orderTable.UserID,
			orderTable.TotalAmount,
			orderTable.Status,
			orderTable.CreatedAt,
			jsonAggQuery,
		).From(orderTableName).
		Join(
			fmt.Sprintf("%s ON %s = %s",
				orderItemTableName,
				orderTableName+"."+orderTable.ID,
				orderItemTableName+"."+orderItemTable.OrderID),
		).
		Where(
			sq.Eq{orderTable.UserID: userID},
		).
		OrderBy(fmt.Sprintf("%s DESC", orderTable.CreatedAt)).
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		GroupBy(orderTableName + "." + orderTable.ID).
		PlaceholderFormat(sq.Dollar)

	if !filter.StartTime.IsZero() {
		query = query.Where(sq.GtOrEq{orderTable.CreatedAt: filter.StartTime})
	}

	if !filter.EndTime.IsZero() {
		query = query.Where(sq.LtOrEq{orderTable.CreatedAt: filter.EndTime})
	}

	res, err := query.RunWith(r.psql.DB()).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		orderDetail := model.OrderDetails{}
		var orderItemStr string

		err = res.Scan(
			&orderDetail.ID,
			&orderDetail.UserID,
			&orderDetail.Total,
			&orderDetail.Status,
			&orderDetail.CreatedAt,
			&orderItemStr,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(orderItemStr), &orderDetail.Items)
		if err != nil {
			return nil, err
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, nil
}
