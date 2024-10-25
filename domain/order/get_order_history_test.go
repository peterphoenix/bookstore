package order

import (
	"context"
	"fmt"
	"testing"

	"bookstore/common/pagination"
	"bookstore/domain/order/model"
	order_mock "bookstore/test/mock/domain/order"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetOrderHistory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := order_mock.NewMockRepository(mockCtrl)

	svc := NewService(mockRepo)

	ctx := context.TODO()
	dummyError := "dummy error"
	getOrderHistoryRes := []model.OrderDetails{
		{
			Order: model.Order{
				ID:     "1",
				UserID: "123",
				Status: "COMPLETED",
			},
			Items: []model.OrderItemDetails{
				{
					BookID: "123",
					Qty:    1,
				},
			},
		},
	}

	t.Run("error get order history", func(t *testing.T) {
		mockRepo.EXPECT().GetOrderHistory(ctx, gomock.Any(), gomock.Any()).
			Return(nil, fmt.Errorf(dummyError))

		res, err := svc.GetOrderHistory(ctx, model.GetOrderHistoryFilter{}, pagination.Pagination{})

		assert.Empty(nil, res)
		assert.EqualError(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetOrderHistory(ctx, gomock.Any(), gomock.Any()).
			Return(getOrderHistoryRes, nil)

		res, err := svc.GetOrderHistory(ctx, model.GetOrderHistoryFilter{}, pagination.Pagination{})

		assert.Nil(t, err)
		assert.Equal(t, getOrderHistoryRes, res)
	})
}
