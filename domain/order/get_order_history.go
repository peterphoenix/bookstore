package order

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/order/model"
)

func (s service) GetOrderHistory(ctx context.Context, filter model.GetOrderHistoryFilter, pagination pagination.Pagination) ([]model.OrderDetails, error) {
	res, err := s.repo.GetOrderHistory(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	return res, nil
}
