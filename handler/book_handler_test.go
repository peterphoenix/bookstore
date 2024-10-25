package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore/common/pagination"
	"bookstore/common/respond"
	"bookstore/domain/book/model"
	commonmodel "bookstore/domain/common/model"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ListBooks(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler, mocks := newTestHandler(mockCtrl)

	dummyError := "dummy error"
	res := model.ListBooksRes{
		Books: []commonmodel.Book{
			{
				ID:    "1",
				Title: "book 1",
				Price: decimal.NewFromInt(10000),
			},
			{
				ID:    "2",
				Title: "book 2",
				Price: decimal.NewFromInt(10000),
			},
		},
	}

	t.Run("error list books", func(t *testing.T) {
		mocks.bookMock.EXPECT().ListBooks(gomock.Any(), pagination.Pagination{
			Limit: 1,
			Page:  2,
		}).
			Return(model.ListBooksRes{}, fmt.Errorf(dummyError))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/book?limit=1&page=2", nil)
		testHandler.ListBooks(w, r)

		var resp respond.APIModel[model.ListBooksRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, respond.CodeInternalServerError, resp.Errors[0].Code)
		assert.Equal(t, dummyError, resp.Errors[0].Message)
	})

	t.Run("success", func(t *testing.T) {
		mocks.bookMock.EXPECT().ListBooks(gomock.Any(), gomock.Any()).
			Return(res, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/book", nil)
		testHandler.ListBooks(w, r)

		var resp respond.APIModel[model.ListBooksRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Empty(t, resp.Errors)
		assert.Equal(t, res, resp.Data)
	})
}
