package handler

import (
	"testing"

	book_mock "bookstore/test/mock/domain/book"
	order_mock "bookstore/test/mock/domain/order"
	user_mock "bookstore/test/mock/domain/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mockStruct struct {
	userMock  *user_mock.MockService
	bookMock  *book_mock.MockService
	orderMock *order_mock.MockService
}

func newTestHandler(mockCtrl *gomock.Controller) (Handler, mockStruct) {
	userSvc := user_mock.NewMockService(mockCtrl)
	bookSvc := book_mock.NewMockService(mockCtrl)
	orderSvc := order_mock.NewMockService(mockCtrl)

	return NewHandler(userSvc, bookSvc, orderSvc), mockStruct{
		userMock:  userSvc,
		bookMock:  bookSvc,
		orderMock: orderSvc}
}

func TestNewHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("success", func(t *testing.T) {
		testHandler, mocks := newTestHandler(mockCtrl)
		assert.NotEmpty(t, testHandler)
		assert.NotEmpty(t, mocks)
	})
}
