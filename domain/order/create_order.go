package order

import (
	"context"

	"bookstore/domain/order/model"
	"github.com/google/uuid"
)

func (s service) CreateOrder(ctx context.Context, input model.CreateOrderInput) (model.CreateOrderRes, error) {
	id := uuid.NewString()
	input.ID = id

	err := s.repo.CreateOrder(ctx, input)
	if err != nil {
		return model.CreateOrderRes{}, err
	}

	return model.CreateOrderRes{
		ID: id,
	}, nil
}
