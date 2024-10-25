package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"bookstore/common/contextutil"
	"bookstore/common/header"
	"bookstore/common/pagination"
	"bookstore/common/respond"
	"bookstore/common/timeutil"
	bookmodel "bookstore/domain/book/model"
	"bookstore/domain/order/model"
	"github.com/shopspring/decimal"
)

// CreateOrder godoc
// @Summary		Create order
// @Description	Endpoint used to create order
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		X-User-ID header string true "X-User-ID"
// @Param		request body model.CreateOrderReq true "body"
// @Success		200 {object} 	respond.APIModel[model.CreateOrderRes]
// @Failure		400	{object}	respond.APIModel[string]
// @Failure		401	{object}	respond.APIModel[string]
// @Failure		500 {object} 	respond.APIModel[string]
// @Router		/order [post]
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Header.Get(header.UserIDKey)
	if len(userId) == 0 {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeUnauthorized,
			Desc: respond.CodeUnauthorized,
		}, http.StatusUnauthorized)
	}
	ctx = contextutil.WithUserID(ctx, userId)

	var input model.CreateOrderReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInvalidRequest,
			Desc: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = input.Validate()
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInvalidRequest,
			Desc: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	var ids []string
	for _, item := range input.Items {
		ids = append(ids, item.BookID)
	}
	books, err := h.BookSvc.GetBookDetailsMap(ctx, bookmodel.GetBookDetailsReq{IDs: ids})
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInternalServerError,
			Desc: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	var items []model.OrderItemDetails
	total := decimal.NewFromInt(0)
	for _, item := range input.Items {
		price := books[item.BookID].Price
		items = append(items, model.OrderItemDetails{
			BookID: item.BookID,
			Qty:    item.Qty,
			Price:  price,
		})
		total = total.Add(price.Mul(decimal.NewFromUint64(uint64(item.Qty))))
	}

	createOrderInput := model.CreateOrderInput{
		Items:           items,
		UserID:          userId,
		Total:           total,
		ShippingAddress: input.ShippingAddress,
		City:            input.City,
		State:           input.State,
		ZipCode:         input.ZipCode,
		Country:         input.Country,
	}

	res, err := h.OrderSvc.CreateOrder(ctx, createOrderInput)
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInternalServerError,
			Desc: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	respond.Success(w, res, http.StatusCreated)
}

// GetOrderHistory godoc
// @Summary		Get order history
// @Description	Endpoint used to get order history
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		X-User-ID header string true "X-User-ID"
// @Param 		limit query string false "limit per page"
// @Param 		page query string false "page number"
// @Param 		startTime query string false "start timestamp, format: 2006-01-01T00:00:00Z"
// @Param 		endTime query string false "end timestamp, format: 2006-01-01T00:00:00Z"
// @Success		200 {object} 	respond.APIModel[model.GetOrderHistoryRes]
// @Failure		400	{object}	respond.APIModel[string]
// @Failure		401	{object}	respond.APIModel[string]
// @Failure		500 {object} 	respond.APIModel[string]
// @Router		/order/history [get]
func (h *handler) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Header.Get(header.UserIDKey)
	if len(userId) == 0 {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeUnauthorized,
			Desc: respond.CodeUnauthorized,
		}, http.StatusUnauthorized)

		return
	}
	ctx = contextutil.WithUserID(ctx, userId)

	startTime := time.Time{}
	endTime := time.Time{}

	startTimeStr := r.URL.Query().Get("startTime")
	if startTimeStr != "" {
		t, err := timeutil.ParseTime(startTimeStr)
		startTime = t
		if err != nil {
			respond.Error(w, ctx, respond.APIError{
				Code: respond.CodeInvalidRequest,
				Desc: err.Error(),
			}, http.StatusBadRequest)

			return
		}
	}

	endTimeStr := r.URL.Query().Get("endTime")
	if endTimeStr != "" {
		t, err := timeutil.ParseTime(endTimeStr)
		endTime = t
		if err != nil {
			respond.Error(w, ctx, respond.APIError{
				Code: respond.CodeInvalidRequest,
				Desc: err.Error(),
			}, http.StatusBadRequest)

			return
		}
	}

	orderDetails, err := h.OrderSvc.GetOrderHistory(ctx, model.GetOrderHistoryFilter{
		StartTime: startTime,
		EndTime:   endTime,
	}, pagination.GetPagination(r))
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInternalServerError,
			Desc: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	var ids []string
	for _, orderDetail := range orderDetails {
		for _, item := range orderDetail.Items {
			ids = append(ids, item.BookID)
		}
	}
	books, err := h.BookSvc.GetBookDetailsMap(ctx, bookmodel.GetBookDetailsReq{IDs: ids})
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInternalServerError,
			Desc: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	var res model.GetOrderHistoryRes

	for _, orderDetail := range orderDetails {
		var items []model.OrderHistoryItemDetails
		for _, item := range orderDetail.Items {
			book := books[item.BookID]
			items = append(items, model.OrderHistoryItemDetails{
				BookID:   item.BookID,
				Qty:      item.Qty,
				Price:    item.Price,
				Title:    book.Title,
				ImageURL: book.ImageURL,
			})
		}

		res.OrderHistoryDetails = append(res.OrderHistoryDetails, model.OrderHistoryDetail{
			ID:        orderDetail.ID,
			UserID:    orderDetail.UserID,
			Status:    orderDetail.Status,
			Total:     orderDetail.Total,
			CreatedAt: orderDetail.CreatedAt,
			Items:     items,
		})
	}

	respond.Success(w, res, http.StatusOK)
}
