package book

import (
	"context"
	"fmt"
	"testing"

	"bookstore/common/pagination"
	commonmodel "bookstore/domain/common/model"
	book_mock "bookstore/test/mock/domain/book"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_ListBooks(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := book_mock.NewMockRepository(mockCtrl)

	svc := NewService(mockRepo)

	ctx := context.TODO()
	dummyError := "dummy error"
	booksListRes := []commonmodel.Book{
		{
			ID:    "1",
			Title: "book 1",
		},
		{
			ID:    "2",
			Title: "book 2",
		},
		{
			ID:    "3",
			Title: "book 3",
		},
	}

	t.Run("error get book list", func(t *testing.T) {
		mockRepo.EXPECT().GetBooksList(ctx, gomock.Any()).
			Return(nil, fmt.Errorf(dummyError))

		res, err := svc.ListBooks(ctx, pagination.Pagination{})

		assert.Empty(t, res)
		assert.EqualError(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetBooksList(ctx, gomock.Any()).
			Return(booksListRes, nil)

		res, err := svc.ListBooks(ctx, pagination.Pagination{})

		assert.Nil(t, err)
		assert.Equal(t, res.Books, booksListRes)
	})
}
