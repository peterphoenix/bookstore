package order

import (
	"context"
	"fmt"
	"testing"

	"bookstore/domain/order/model"
	order_mock "bookstore/test/mock/domain/order"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := order_mock.NewMockRepository(mockCtrl)

	svc := NewService(mockRepo)

	ctx := context.TODO()
	dummyError := "dummy error"

	t.Run("error create order", func(t *testing.T) {
		mockRepo.EXPECT().CreateOrder(ctx, gomock.Any()).
			Return(fmt.Errorf(dummyError))

		res, err := svc.CreateOrder(ctx, model.CreateOrderInput{})

		assert.Empty(t, res)
		assert.EqualError(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().CreateOrder(ctx, gomock.Any()).
			Return(nil)

		res, err := svc.CreateOrder(ctx, model.CreateOrderInput{})

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})
}
