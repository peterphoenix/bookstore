package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bookstore/common/respond"
	bookmodel "bookstore/domain/book/model"
	"bookstore/domain/order/model"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler, mocks := newTestHandler(mockCtrl)

	dummyError := "dummy error"
	getInput := func() model.CreateOrderReq {
		return model.CreateOrderReq{
			Items: []model.CreateOrderItemReq{
				{
					BookID: "123",
					Qty:    1,
				},
			},
		}
	}

	bookDetails := bookmodel.BookDetailsMap{
		"123": bookmodel.BookDetails{
			ID:    "123",
			Title: "book 1",
		},
	}

	res := model.CreateOrderRes{
		ID: "123",
	}

	t.Run("error missing user id", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", nil)
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, respond.CodeUnauthorized, resp.Errors[0].Code)
		assert.NotEmpty(t, resp.Errors[0].Message)
	})

	t.Run("error decode input", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", nil)
		r.Header.Set("X-User-ID", "123")
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.NotEmpty(t, resp.Errors[0].Message)
	})

	t.Run("error empty item", func(t *testing.T) {
		input := getInput()
		input.Items = nil
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(inputJSON))
		r.Header.Set("X-User-ID", "123")
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.Equal(t, model.EmptyItemErr, resp.Errors[0].Message)
	})

	t.Run("error zero qty", func(t *testing.T) {
		input := getInput()
		input.Items[0].Qty = 0
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(inputJSON))
		r.Header.Set("X-User-ID", "123")
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.Equal(t, model.ZeroQtyErr, resp.Errors[0].Message)
	})

	t.Run("error get book details", func(t *testing.T) {
		input := getInput()
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		mocks.bookMock.EXPECT().GetBookDetailsMap(gomock.Any(), gomock.Any()).
			Return(nil, fmt.Errorf(dummyError))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(inputJSON))
		r.Header.Set("X-User-ID", "123")
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, respond.CodeInternalServerError, resp.Errors[0].Code)
		assert.Equal(t, dummyError, resp.Errors[0].Message)
	})

	t.Run("error create order", func(t *testing.T) {
		input := getInput()
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		mocks.bookMock.EXPECT().GetBookDetailsMap(gomock.Any(), gomock.Any()).
			Return(bookDetails, nil)
		mocks.orderMock.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).
			Return(model.CreateOrderRes{}, fmt.Errorf(dummyError))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(inputJSON))
		r.Header.Set("X-User-ID", "123")
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, respond.CodeInternalServerError, resp.Errors[0].Code)
		assert.Equal(t, dummyError, resp.Errors[0].Message)
	})

	t.Run("success", func(t *testing.T) {
		input := getInput()
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		mocks.bookMock.EXPECT().GetBookDetailsMap(gomock.Any(), gomock.Any()).
			Return(bookDetails, nil)
		mocks.orderMock.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).
			Return(res, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(inputJSON))
		r.Header.Set("X-User-ID", "123")
		testHandler.CreateOrder(w, r)

		var resp respond.APIModel[model.CreateOrderRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Empty(t, resp.Errors)
		assert.Equal(t, res, resp.Data)
	})
}

func TestHandler_GetOrderHistory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler, mocks := newTestHandler(mockCtrl)

	dummyError := "dummy error"

	orderHistory := []model.OrderDetails{
		{
			Order: model.Order{
				ID:        "123",
				UserID:    "123",
				Status:    "COMPLETED",
				Total:     decimal.NewFromInt(10000),
				CreatedAt: time.Time{},
			},
			Items: []model.OrderItemDetails{
				{
					BookID: "123",
					Qty:    1,
					Price:  decimal.NewFromInt(10000),
				},
			},
		},
	}
	bookDetails := bookmodel.BookDetailsMap{
		"123": bookmodel.BookDetails{
			ID:       "123",
			Title:    "book 1",
			ImageURL: "www.example.com",
		},
	}

	res := model.GetOrderHistoryRes{
		OrderHistoryDetails: []model.OrderHistoryDetail{
			{
				ID:        "123",
				UserID:    "123",
				Status:    "COMPLETED",
				Total:     decimal.NewFromInt(10000),
				CreatedAt: time.Time{},
				Items: []model.OrderHistoryItemDetails{
					{
						BookID:   "123",
						Qty:      1,
						Price:    decimal.NewFromInt(10000),
						Title:    "book 1",
						ImageURL: "www.example.com",
					},
				},
			},
		},
	}

	t.Run("error missing user id", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order/history", nil)
		testHandler.GetOrderHistory(w, r)

		var resp respond.APIModel[model.GetOrderHistoryRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, respond.CodeUnauthorized, resp.Errors[0].Code)
		assert.NotEmpty(t, resp.Errors[0].Message)
	})

	t.Run("error parse start time", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order/history?startTime=123", nil)
		r.Header.Set("X-User-ID", "123")
		testHandler.GetOrderHistory(w, r)

		var resp respond.APIModel[model.GetOrderHistoryRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.NotEmpty(t, resp.Errors[0].Message)
	})

	t.Run("error parse end time", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/order/history?endTime=123", nil)
		r.Header.Set("X-User-ID", "123")
		testHandler.GetOrderHistory(w, r)

		var resp respond.APIModel[model.GetOrderHistoryRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.NotEmpty(t, resp.Errors[0].Message)
	})

	t.Run("error get order history", func(t *testing.T) {
		mocks.orderMock.EXPECT().GetOrderHistory(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, fmt.Errorf(dummyError))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/order/history", nil)
		r.Header.Set("X-User-ID", "123")
		testHandler.GetOrderHistory(w, r)

		var resp respond.APIModel[model.GetOrderHistoryRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, respond.CodeInternalServerError, resp.Errors[0].Code)
		assert.Equal(t, dummyError, resp.Errors[0].Message)
	})

	t.Run("error get book details", func(t *testing.T) {
		mocks.orderMock.EXPECT().GetOrderHistory(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(orderHistory, nil)
		mocks.bookMock.EXPECT().GetBookDetailsMap(gomock.Any(), gomock.Any()).
			Return(nil, fmt.Errorf(dummyError))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/order/history", nil)
		r.Header.Set("X-User-ID", "123")
		testHandler.GetOrderHistory(w, r)

		var resp respond.APIModel[model.GetOrderHistoryRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, respond.CodeInternalServerError, resp.Errors[0].Code)
		assert.Equal(t, dummyError, resp.Errors[0].Message)
	})

	t.Run("success", func(t *testing.T) {
		mocks.orderMock.EXPECT().GetOrderHistory(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(orderHistory, nil)
		mocks.bookMock.EXPECT().GetBookDetailsMap(gomock.Any(), gomock.Any()).
			Return(bookDetails, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/order/history?startTime=2024-01-01T00:00:00Z&endTime=2034-01-01T00:00:00Z", nil)
		r.Header.Set("X-User-ID", "123")
		testHandler.GetOrderHistory(w, r)

		var resp respond.APIModel[model.GetOrderHistoryRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, resp.Errors)
		assert.Equal(t, res, resp.Data)
	})
}
