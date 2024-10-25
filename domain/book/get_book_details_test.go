package book

import (
	"context"
	"fmt"
	"testing"

	"bookstore/domain/book/model"
	book_mock "bookstore/test/mock/domain/book"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetBookDetailsMap(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := book_mock.NewMockRepository(mockCtrl)

	svc := NewService(mockRepo)

	ctx := context.TODO()
	dummyError := "dummy error"
	ids := []string{"1", "2", "3"}
	bookDetailsRes := []model.BookDetails{
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

	t.Run("error get book details", func(t *testing.T) {
		mockRepo.EXPECT().GetBookDetails(ctx, ids).
			Return(nil, fmt.Errorf(dummyError))

		res, err := svc.GetBookDetailsMap(ctx, model.GetBookDetailsReq{IDs: ids})

		assert.Nil(t, res)
		assert.EqualError(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().GetBookDetails(ctx, ids).
			Return(bookDetailsRes, nil)

		res, err := svc.GetBookDetailsMap(ctx, model.GetBookDetailsReq{IDs: ids})

		assert.Nil(t, err)
		assert.Equal(t, bookDetailsRes[0], res["1"])
		assert.Equal(t, bookDetailsRes[1], res["2"])
		assert.Equal(t, bookDetailsRes[2], res["3"])
	})
}
