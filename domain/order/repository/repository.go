package repository

//go:generate mockgen -source ./repository.go -destination ../../../test/mock/domain/order/repository.go -package=order_mock

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/order/model"
	"bookstore/infra/database"
)

type Repository interface {
	CreateOrder(ctx context.Context, input model.CreateOrderInput) error
	GetOrderHistory(ctx context.Context, filter model.GetOrderHistoryFilter, pagination pagination.Pagination) ([]model.OrderDetails, error)
}

type repository struct {
	psql *database.PostgreSQL
}

func NewRepository(psql *database.PostgreSQL) Repository {
	return &repository{psql}
}
