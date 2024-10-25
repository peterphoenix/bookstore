package order

//go:generate mockgen -source ./service.go -destination ../../test/mock/domain/order/service.go -package=order_mock

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/order/model"
	"bookstore/domain/order/repository"
)

type Service interface {
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (model.CreateOrderRes, error)
	GetOrderHistory(ctx context.Context, filter model.GetOrderHistoryFilter, pagination pagination.Pagination) ([]model.OrderDetails, error)
}

type service struct {
	repo repository.Repository
}

func NewService(
	repo repository.Repository,
) Service {
	return &service{
		repo: repo,
	}
}
