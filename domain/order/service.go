package order

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/order/model"
	"bookstore/domain/order/repository"
	"bookstore/infra/database"
)

type Service interface {
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (model.CreateOrderRes, error)
	GetOrderHistory(ctx context.Context, filter model.GetOrderHistoryFilter, pagination pagination.Pagination) ([]model.OrderDetails, error)
}

type service struct {
	repo repository.Repository
}

func NewService(
	db *database.PostgreSQL,
) Service {
	return &service{
		repo: repository.NewRepository(db),
	}
}
